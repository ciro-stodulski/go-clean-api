package notificationservice

import (
	domaindto "go-clean-api/cmd/domain/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ns notificationService) FindById(msg string) *domaindto.Event {

	id, _ := primitive.ObjectIDFromHex(msg)
	result, err := ns.NotificationCollection.FindById(id)

	if err != nil {
		println(err)
	}

	return result
}
