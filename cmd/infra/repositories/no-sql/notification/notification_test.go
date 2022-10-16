package notificationcollection

import (
	portsservice "go-clean-api/cmd/core/ports"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestFindById(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	defer mt.Close()

	mt.Run("should return notification with success", func(mt *mtest.T) {
		mt.DB = mt.Client.Database("go-clean-api")
		notification_fake := portsservice.Dto{
			Name:  "john",
			Event: "created",
		}

		id_fake := primitive.NewObjectID()

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "db.notifications", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: id_fake},
			{Key: "name", Value: notification_fake.Name},
			{Key: "event", Value: notification_fake.Event},
		}))

		collection := New(mt.DB)

		result, err := collection.FindById(id_fake)

		assert.Nil(t, err)
		assert.Equal(t, notification_fake, portsservice.Dto{
			Name:  result.Name,
			Event: result.Event,
		})
	})
}

func TestCreate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	defer mt.Close()

	mt.Run("should create notification with success", func(mt *mtest.T) {
		mt.DB = mt.Client.Database("go-clean-api")

		notification_fake := portsservice.Dto{
			Name:  "john",
			Event: "created",
		}
		id_fake := primitive.NewObjectID()

		mt.AddMockResponses(mtest.CreateSuccessResponse(
			bson.D{
				{Key: "_id", Value: id_fake},
				{Key: "name", Value: notification_fake.Name},
				{Key: "event", Value: notification_fake.Event},
			}...,
		))

		collection := New(mt.DB)

		result := collection.Create(notification_fake)

		assert.NotNil(t, result)
	})
}