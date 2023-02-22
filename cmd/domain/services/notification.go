package domainservice

import domaindto "go-clean-api/cmd/domain/dto"

type (
	NotificationService interface {
		SendNotify(dto domaindto.Event) error
		CheckNotify(msg string) (string error)
		SaveNotify(domaindto.Event) string
		FindById(id string) *domaindto.Event
	}
)
