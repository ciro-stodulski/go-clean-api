package domainservice

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	domainexceptions "go-clean-api/cmd/domain/exceptions"

	"go-clean-api/cmd/domain/entities/user"
)

type (
	UserService interface {
		GetUser(id string) (*user.User, *domainexceptions.ApplicationException, error)
		GetByEmail(email string) (*user.User, *domainexceptions.ApplicationException, error)
		Register(*user.User) (*user.User, *domainexceptions.ApplicationException, error)
		DeleteUser(id string) (*domainexceptions.ApplicationException, error)
		ListUsers() ([]response_jsonplaceholder.User, *domainexceptions.ApplicationException, error)
	}
)
