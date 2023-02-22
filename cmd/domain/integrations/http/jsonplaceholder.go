package domainjsonplaceholder

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
)

type (
	JsonPlaceholderIntegration interface {
		GetUsers() ([]response_jsonplaceholder.User, error)
	}
)
