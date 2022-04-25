package system

import (
	"api/app/departments"
	"api/app/pages"
	"api/app/roles"
	"api/app/users"
	"api/common"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thoas/go-funk"
	"github.com/weplanx/go/helper"
	"github.com/weplanx/go/passport"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Controller struct {
	Service     *Service
	Users       *users.Service
	Roles       *roles.Service
	Departments *departments.Service
	Pages       *pages.Service
	Passport    *passport.Passport
}

func (x *Controller) Index(c *gin.Context) interface{} {
	return gin.H{
		"time": time.Now(),
	}
}

// AuthLogin 登录
func (x *Controller) AuthLogin(c *gin.Context) interface{} {
	var body struct {
		User     string `json:"user" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	ctx := c.Request.Context()
	// 用户验证
	data, err := x.Users.FindOneByUsernameOrEmail(ctx, body.User)
	if err != nil {
		c.Set("code", "AUTH_INCORRECT")
		return err
	}
	if err := helper.PasswordVerify(body.Password, data.Password); err != nil {
		c.Set("code", "AUTH_INCORRECT")
		return err
	}
	// 创建 Token
	jti := helper.Uuid()
	var ts string
	if ts, err = x.Passport.Create(jti, gin.H{
		"uid": data.ID.Hex(),
	}); err != nil {
		return err
	}
	// 设置会话
	if err := x.Service.SetSession(ctx, data.ID.Hex(), jti); err != nil {
		return err
	}
	// 写入日志
	dto := common.NewLoginLogV10(data, jti, c.ClientIP(), c.Request.UserAgent())
	go x.Service.WriteLoginLog(context.TODO(), dto)
	// 返回
	c.SetCookie("access_token", ts, 0, "", "", true, true)
	c.SetSameSite(http.SameSiteStrictMode)
	return nil
}

// AuthVerify 主动验证
func (x *Controller) AuthVerify(c *gin.Context) interface{} {
	ts, err := c.Cookie("access_token")
	if err != nil {
		c.Set("status_code", 401)
		c.Set("code", "AUTH_EXPIRED")
		return common.AuthExpired
	}
	if _, err = x.Passport.Verify(ts); err != nil {
		c.Set("status_code", 401)
		c.Set("code", "AUTH_EXPIRED")
		return err
	}
	return nil
}

// AuthCode 申请刷新验证码
func (x *Controller) AuthCode(c *gin.Context) interface{} {
	claims, exists := c.Get(common.TokenClaimsKey)
	if !exists {
		c.Set("status_code", 401)
		c.Set("code", "AUTH_EXPIRED")
		return common.AuthExpired
	}
	jti := claims.(jwt.MapClaims)["jti"].(string)
	code := funk.RandomString(8)
	ctx := c.Request.Context()
	if err := x.Service.CreateVerifyCode(ctx, jti, code); err != nil {
		return err
	}
	return gin.H{"code": code}
}

// AuthRefresh 刷新认证
func (x *Controller) AuthRefresh(c *gin.Context) interface{} {
	var body struct {
		Code string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	// 获取载荷
	value, exists := c.Get(common.TokenClaimsKey)
	if !exists {
		c.Set("status_code", 401)
		c.Set("code", "AUTH_EXPIRED")
		return common.AuthExpired
	}
	claims := value.(jwt.MapClaims)
	jti := claims["jti"].(string)
	ctx := c.Request.Context()
	// 刷新验证
	result, err := x.Service.VerifyCode(ctx, jti, body.Code)
	if err != nil {
		return err
	}
	if !result {
		c.Set("status_code", 401)
		c.Set("code", "AUTH_EXPIRED")
		return common.AuthExpired

	}
	if err = x.Service.DeleteVerifyCode(ctx, jti); err != nil {
		return err
	}
	// 继承 jti 创建新 Token
	var ts string
	if ts, err = x.Passport.Create(jti,
		claims["context"].(map[string]interface{}),
	); err != nil {
		return err
	}
	c.SetCookie("access_token", ts, 0, "", "", true, true)
	c.SetSameSite(http.SameSiteStrictMode)
	return nil
}

// AuthLogout 登出
func (x *Controller) AuthLogout(c *gin.Context) interface{} {
	c.SetCookie("access_token", "", 0, "", "", true, true)
	c.SetSameSite(http.SameSiteStrictMode)
	return nil
}

// GetUser 获取用户信息
func (x *Controller) GetUser(c *gin.Context) interface{} {
	claims, exists := c.Get(common.TokenClaimsKey)
	if !exists {
		c.Set("status_code", 401)
		c.Set("code", "AUTH_EXPIRED")
		return common.AuthExpired
	}
	ctx := c.Request.Context()
	uid := claims.(jwt.MapClaims)["context"].(map[string]interface{})["uid"].(string)
	var data common.User
	if err := x.Users.FindOneById(ctx, uid, &data); err != nil {
		return err
	}
	result := gin.H{
		"username":    data.Username,
		"email":       data.Email,
		"name":        data.Name,
		"avatar":      data.Avatar,
		"sessions":    data.Sessions,
		"last":        data.Last,
		"create_time": data.CreateTime,
	}
	var err error
	if result["roles"], err = x.Roles.FindNamesById(ctx, data.Roles); err != nil {
		return err
	}
	if data.Department != nil {
		if result["department"], err = x.Departments.FindNameById(ctx, *data.Department); err != nil {
			return err
		}
	}
	return result
}

func (x *Controller) SetUser(c *gin.Context) interface{} {
	var body struct {
		Name   string `json:"name,omitempty"`
		Avatar string `json:"avatar,omitempty"`
		Email  string `json:"email,omitempty"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	claims, exists := c.Get(common.TokenClaimsKey)
	if !exists {
		c.Set("status_code", 401)
		c.Set("code", "AUTH_EXPIRED")
		return common.AuthExpired
	}
	ctx := c.Request.Context()
	uid := claims.(jwt.MapClaims)["context"].(map[string]interface{})["uid"].(string)
	if err := x.Users.UpdateOneById(ctx, uid, bson.M{
		"$set": body,
	}); err != nil {
		return err
	}
	return nil
}

// GetVars 获取指定变量
func (x *Controller) GetVars(c *gin.Context) interface{} {
	var query struct {
		Keys []string `form:"keys" binding:"required"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		return err
	}
	ctx := c.Request.Context()
	values, err := x.Service.GetVars(ctx, query.Keys)
	if err != nil {
		return err
	}
	return values
}

// GetVar 获取变量
func (x *Controller) GetVar(c *gin.Context) interface{} {
	var uri struct {
		Key string `uri:"key" binding:"required"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		return err
	}
	ctx := c.Request.Context()
	value, err := x.Service.GetVar(ctx, uri.Key)
	if err != nil {
		return err
	}
	return value
}

// SetVar 设置变量
func (x *Controller) SetVar(c *gin.Context) interface{} {
	var uri struct {
		Key string `uri:"key" binding:"required"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		return err
	}
	var body struct {
		Value interface{} `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	ctx := c.Request.Context()
	if err := x.Service.SetVar(ctx, uri.Key, body.Value); err != nil {
		return err
	}
	return nil
}

// GetSessions 获取会话
func (x *Controller) GetSessions(c *gin.Context) interface{} {
	ctx := c.Request.Context()
	values, err := x.Service.GetSessions(ctx)
	if err != nil {
		return err
	}
	return values
}

// DeleteSessions 删除所有会话
func (x *Controller) DeleteSessions(c *gin.Context) interface{} {
	ctx := c.Request.Context()
	if err := x.Service.DeleteSessions(ctx); err != nil {
		return err
	}
	return nil
}

// DeleteSession 删除会话
func (x *Controller) DeleteSession(c *gin.Context) interface{} {
	var uri struct {
		Id string `uri:"id" binding:"required,objectId"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		return err
	}
	ctx := c.Request.Context()
	if err := x.Service.DeleteSession(ctx, uri.Id); err != nil {
		return err
	}
	return nil
}

// Sort 通用排序
func (x *Controller) Sort(c *gin.Context) interface{} {
	var uri struct {
		Model string `uri:"model"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		return err
	}
	var body struct {
		Sort []primitive.ObjectID `json:"sort" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}
	ctx := c.Request.Context()
	result, err := x.Service.Sort(ctx, uri.Model, body.Sort)
	if err != nil {
		return err
	}
	return result
}

// Uploader 上传签名
func (x *Controller) Uploader(c *gin.Context) interface{} {
	data, err := x.Service.Uploader()
	if err != nil {
		return err
	}
	return data
}

// Navs 页面导航
func (x *Controller) Navs(c *gin.Context) interface{} {
	ctx := c.Request.Context()
	navs, err := x.Pages.Navs(ctx)
	if err != nil {
		return err
	}
	return navs
}

func (x *Controller) Dynamic(c *gin.Context) interface{} {
	var uri struct {
		Id string `uri:"id" binding:"required,objectId"`
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		return err
	}
	ctx := c.Request.Context()
	data, err := x.Pages.FindOneById(ctx, uri.Id)
	if err != nil {
		return err
	}
	return data
}
