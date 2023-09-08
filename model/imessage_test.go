package model_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
)

func TestCreateImessages(t *testing.T) {
	ctx := context.TODO()
	b, err := os.ReadFile("./model/imessage.json")
	assert.NoError(t, err)
	var jsonSchema bson.D
	err = bson.UnmarshalExtJSON(b, true, &jsonSchema)
	assert.NoError(t, err)

	n, err := x.Db.ListCollectionNames(ctx, bson.M{"name": "imessages"})
	assert.NoError(t, err)
	if len(n) == 0 {
		option := options.CreateCollection().SetValidator(jsonSchema)
		err = x.Db.CreateCollection(ctx, "imessages", option)
		assert.NoError(t, err)
	} else {
		err = x.Db.RunCommand(ctx, bson.D{
			{"collMod", "imessages"},
			{"validator", jsonSchema},
			{"validationLevel", "strict"},
		}).Err()
		assert.NoError(t, err)
	}
}