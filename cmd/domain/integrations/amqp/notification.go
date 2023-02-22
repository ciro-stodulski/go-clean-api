package domainnotificationproducer

import (
	portsservice "go-clean-api/cmd/domain/services"
)

type (
	NotificationProducer interface {
		SendNotify(dto portsservice.Dto) error
	}
)
