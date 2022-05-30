package ports

import (
	integration "go-api/src/infra/integrations/http/jsonplaceholder/responses"
)

type (
	JsonPlaceholderIntegration interface {
		GetUsers() ([]integration.User, error)
	}
)
