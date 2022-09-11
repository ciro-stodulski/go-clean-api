package container

import (
	database "go-clean-api/cmd/infra/adapters/mysql"
	amqpclient "go-clean-api/cmd/infra/integrations/amqp"
	grpc_client "go-clean-api/cmd/infra/integrations/grpc"
	http_service "go-clean-api/cmd/infra/integrations/http"
	cache_client "go-clean-api/cmd/infra/repositories/cache"

	"github.com/jinzhu/gorm"
)

type (
	containerConfig struct {
		Database     *gorm.DB
		Grpc_client  grpc_client.GRPCClient
		Amqp_client  amqpclient.AmqpClient
		Http_client  http_service.HttpClient
		Cache_client cache_client.CacheClient
	}
)

var db database.MysqlAdapter

func newContainerConfig() containerConfig {
	db.ConnectToDatabase()

	return containerConfig{
		Database:     db.Db,
		Grpc_client:  grpc_client.New(),
		Amqp_client:  amqpclient.New(),
		Http_client:  http_service.New(),
		Cache_client: cache_client.New(),
	}
}
