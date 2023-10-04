package container

import (
	messagingentity "go-clean-api/cmd/domain/entity/messaging"
	httpadapter "go-clean-api/cmd/infra/adapters/http"
	adaptermongodb "go-clean-api/cmd/infra/adapters/mongodb"
	database "go-clean-api/cmd/infra/adapters/mysql"
	redisadapter "go-clean-api/cmd/infra/adapters/redis"
	amqpclient "go-clean-api/cmd/infra/integration/amqp"
	grpc_client "go-clean-api/cmd/infra/integration/grpc"
	http_service "go-clean-api/cmd/infra/integration/http"
	cache_client "go-clean-api/cmd/infra/repository/cache"
	"go-clean-api/cmd/presentation/http/controller"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	containerConfig struct {
		Database          *gorm.DB
		DatabaseNoSql     *mongo.Database
		GrpcClient        grpc_client.GRPCClient
		AmqpClient        amqpclient.AmqpClient
		HttpClient        http_service.HttpClient
		CacheClient       cache_client.CacheClient
		SubjectIDChannels map[string]controller.ChannelManager[messagingentity.MessagingEntity]
	}
)

var db database.MysqlAdapter

func newContainerConfig() containerConfig {
	db.ConnectToDatabase()
	subjectIDChannels := make(map[string]controller.ChannelManager[messagingentity.MessagingEntity])

	return containerConfig{
		SubjectIDChannels: subjectIDChannels,
		Database:          db.DB,
		DatabaseNoSql:     adaptermongodb.GetClient(),
		GrpcClient:        grpc_client.New(),
		AmqpClient:        amqpclient.New(),
		HttpClient:        httpadapter.New(),
		CacheClient:       redisadapter.New(),
	}
}
