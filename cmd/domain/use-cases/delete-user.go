package domainusecases

import domainexceptions "go-clean-api/cmd/domain/exceptions"

type (
	DeleteUserUseCase interface {
		DeleteUser(id string) (*domainexceptions.ApplicationException, error)
	}
)
