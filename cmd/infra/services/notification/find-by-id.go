package notificationservice

import (
	portsservice "go-clean-api/cmd/domain/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ns notificationService) FindById(msg string) *portsservice.Dto {

	id, _ := primitive.ObjectIDFromHex(msg)
	result, err := ns.NotificationCollection.FindById(id)

	if err != nil {
		println(err)
	}

	return result
}
