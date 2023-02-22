package domainnotificationcollection

import (
	portsservice "go-clean-api/cmd/domain/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	NotificationCollection interface {
		FindById(id primitive.ObjectID) (*portsservice.Dto, error)
		Create(notification portsservice.Dto) string
	}
)
