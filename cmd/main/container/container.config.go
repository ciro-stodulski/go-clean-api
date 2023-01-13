package container

import (
	database "go-clean-api/cmd/infra/adapters/mysql"
	http_service "go-clean-api/cmd/infra/integrations/http"
)

type (
	containerConfig struct {
		Http_client http_service.HttpClient
	}
)

var db database.MysqlAdapter

func newContainerConfig() containerConfig {
	db.ConnectToDatabase()

	return containerConfig{
		Http_client: http_service.New(),
	}
}
