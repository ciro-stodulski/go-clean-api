package domainusecases

import (
	portsservice "go-clean-api/cmd/domain/services"
)

type (
	NotifyUserUseCase interface {
		Notify(dto portsservice.Dto) error
	}
)
