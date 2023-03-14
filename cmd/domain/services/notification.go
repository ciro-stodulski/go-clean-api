package domainservice

import (
	domaindto "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
)

type (
	NotificationService interface {
		SendNotify(dto domaindto.Event) (*domainexceptions.ApplicationException, error)
		CheckNotify(msg string) (*domainexceptions.ApplicationException, error)
		SaveNotify(domaindto.Event) string
		FindById(id string) (*domaindto.Event, *domainexceptions.ApplicationException, error)
	}
)
