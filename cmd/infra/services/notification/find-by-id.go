package notificationservice

import (
	domaindto "go-clean-api/cmd/domain/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ns notificationService) FindById(msg string) (*domaindto.Event, error) {
	id, _ := primitive.ObjectIDFromHex(msg)

	result, err := ns.NotificationCollection.FindById(id)

	return result, err
}
