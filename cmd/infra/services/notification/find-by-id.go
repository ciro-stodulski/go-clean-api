package notificationservice

import (
	domaindto "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ns notificationService) FindById(msg string) (*domaindto.Event, *domainexceptions.ApplicationException, error) {
	id, _ := primitive.ObjectIDFromHex(msg)

	result, err := ns.NotificationCollection.FindById(id)

	return result, nil, err
}
