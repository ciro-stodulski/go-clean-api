package factories

import (
	http_service "go-clean-api/cmd/infra/integrations/http"
	json_place_holder "go-clean-api/cmd/infra/integrations/http/jsonplaceholder"
)

type (
	InfraContext struct {
		Json_place_holder_integration json_place_holder.JsonPlaceholderIntegration
	}
)

func MakeInfraContext(
	http_service http_service.HttpClient,
) InfraContext {
	return InfraContext{
		Json_place_holder_integration: json_place_holder.New(http_service),
	}
}
