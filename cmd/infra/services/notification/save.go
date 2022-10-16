package notificationservice

import (
	portsservice "go-clean-api/cmd/core/ports"
)

func (ns notificationService) SaveNotify(dto portsservice.Dto) string {

	result := ns.NotificationCollection.Create(dto)

	return result
}
