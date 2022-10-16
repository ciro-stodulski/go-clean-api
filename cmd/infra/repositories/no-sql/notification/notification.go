package notificationcollection

import (
	"context"
	portsservice "go-clean-api/cmd/core/ports"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	NotificationCollection interface {
		FindById(id primitive.ObjectID) (*portsservice.Dto, error)
		Create(notification portsservice.Dto) string
	}
	notificationCollection struct {
		collection *mongo.Collection
	}
)

func New(db *mongo.Database) NotificationCollection {
	return &notificationCollection{
		collection: db.Collection("notifications"),
	}
}

func (mc *notificationCollection) FindById(id primitive.ObjectID) (*portsservice.Dto, error) {
	var notification portsservice.Dto

	filter := bson.D{{Key: "_id", Value: id}}

	err := mc.collection.FindOne(context.Background(), filter).Decode(&notification)

	if err != nil {
		log.Default().Println(err)
		return nil, err
	}

	return &notification, nil
}

func (mc *notificationCollection) Create(notification portsservice.Dto) string {
	ctx := context.Background()

	result, err := mc.collection.InsertOne(ctx, notification)

	if err != nil {
		log.Default().Fatal(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex()
}
