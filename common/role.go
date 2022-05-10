package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Role struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`

	// 名称
	Name string `bson:"name" json:"name"`

	// 描述
	Description string `bson:"description" json:"description"`

	// 授权页面
	Pages []primitive.ObjectID `bson:"pages" json:"pages"`

	// 只读权限
	Readonly []primitive.ObjectID `bson:"readonly" json:"readonly"`

	// 标记
	Labels []string `bson:"labels" json:"labels"`

	// 状态
	Status *bool `bson:"status" json:"status"`

	// 创建时间
	CreateTime time.Time `bson:"create_time" json:"-"`

	// 更新时间
	UpdateTime time.Time `bson:"update_time" json:"-"`
}

func NewRole(name string) *Role {
	return &Role{
		Name:       name,
		Pages:      []primitive.ObjectID{},
		Readonly:   []primitive.ObjectID{},
		Labels:     []string{},
		Status:     BoolToP(true),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}

func (x *Role) SetDescription(v string) *Role {
	x.Description = v
	return x
}

func (x *Role) SetPages(v []primitive.ObjectID) *Role {
	x.Pages = v
	return x
}

func (x *Role) SetReadonly(v []primitive.ObjectID) *Role {
	x.Readonly = v
	return x
}

func (x *Role) SetLabel(v string) *Role {
	x.Labels = append(x.Labels, v)
	return x
}