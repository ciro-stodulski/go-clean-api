package domainusecases

import (
	domaindto "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
)

type (
	NotifyUserUseCase interface {
		Notify(dto domaindto.Event) (*domainexceptions.ApplicationException, error)
	}
)
