package domainnotificationproducer

import domaindto "go-clean-api/cmd/domain/dto"

type (
	NotificationProducer interface {
		SendNotify(dto domaindto.Event) error
	}
)
