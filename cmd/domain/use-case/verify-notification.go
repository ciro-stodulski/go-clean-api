package domainusecases

import (
	domaindto "go-clean-api/cmd/domain/dto"
)

type (
	NotifyUserUseCase interface {
		Notify(dto domaindto.Event) error
	}
)
