package notificationservice

import (
	domaindto "go-clean-api/cmd/domain/dto"
)

func (ns notificationService) SaveNotify(dto domaindto.Event) string {

	result := ns.NotificationCollection.Create(dto)

	return result
}
