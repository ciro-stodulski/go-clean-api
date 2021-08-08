package ports

import (
	integration "go-api/src/infra/http/integrations/jsonplaceholder/responses"
)

type (
	JsonPlaceholderIntegration interface {
		GetUsers() ([]integration.User, error)
	}
)
