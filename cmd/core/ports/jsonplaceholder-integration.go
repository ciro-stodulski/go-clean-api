package ports

import (
	integration "go-api/cmd/infra/integrations/http/jsonplaceholder/responses"
)

type (
	JsonPlaceholderIntegration interface {
		GetUsers() ([]integration.User, error)
	}
)
