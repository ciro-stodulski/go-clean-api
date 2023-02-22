package notificationcollection

import (
	"context"
	domaindto "go-clean-api/cmd/domain/dto"
	domainnotificationcollection "go-clean-api/cmd/domain/repositories/no-sql"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	notificationCollection struct {
		collection *mongo.Collection
	}
)

func New(db *mongo.Database) domainnotificationcollection.NotificationCollection {
	return &notificationCollection{
		collection: db.Collection("notifications"),
	}
}

func (mc *notificationCollection) FindById(id primitive.ObjectID) (*domaindto.Event, error) {
	var notification domaindto.Event

	filter := bson.D{{Key: "_id", Value: id}}

	err := mc.collection.FindOne(context.Background(), filter).Decode(&notification)

	if err != nil {
		log.Default().Println(err)
		return nil, err
	}

	return &notification, nil
}

func (mc *notificationCollection) Create(notification domaindto.Event) string {
	ctx := context.Background()

	result, err := mc.collection.InsertOne(ctx, notification)

	if err != nil {
		log.Default().Fatal(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex()
}
