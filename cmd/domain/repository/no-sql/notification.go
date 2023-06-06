package domainnotificationcollection

import (
	domaindto "go-clean-api/cmd/domain/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	NotificationCollection interface {
		FindById(id primitive.ObjectID) (*domaindto.Event, error)
		Create(notification domaindto.Event) string
	}
)
