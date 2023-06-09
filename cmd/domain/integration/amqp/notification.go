package domainnotificationproducer

import "go-clean-api/cmd/domain/dto"

type (
	NotificationProducer interface {
		SendNotify(data dto.Event) error
	}
)
