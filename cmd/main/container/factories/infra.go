package factories

import (
	amqpclient "go-api/cmd/infra/integrations/amqp"
	notificationproducer "go-api/cmd/infra/integrations/amqp/notification"
	grpc_client "go-api/cmd/infra/integrations/grpc"
	http_service "go-api/cmd/infra/integrations/http"
	json_place_holder "go-api/cmd/infra/integrations/http/jsonplaceholder"
	usersjsonplaceholdercache "go-api/cmd/infra/repositories/cache/users-jsonplaceholder"
	usersql "go-api/cmd/infra/repositories/sql/user"

	notificationpbgrpc "go-api/cmd/infra/integrations/grpc/notification"
	"go-api/cmd/infra/integrations/grpc/notification/pb"
	cache_client "go-api/cmd/infra/repositories/cache"
	"go-api/cmd/shared/env"

	"github.com/jinzhu/gorm"
)

type (
	InfraContext struct {
		NotificationPbGrpc            notificationpbgrpc.NotificationPbGrpc
		Notification_amqp             notificationproducer.NotificationProducer
		Json_place_holder_integration json_place_holder.JsonPlaceholderIntegration
		User_repository               usersql.UserSql
		Users_cache                   usersjsonplaceholdercache.UsersJsonPlaceholderCache
	}
)

func MakeInfraContext(
	grpc_client grpc_client.GRPCClient,
	amqp_client amqpclient.AmqpClient,
	http_service http_service.HttpClient,
	database *gorm.DB,
	cache_client cache_client.CacheClient,
) InfraContext {
	return InfraContext{
		NotificationPbGrpc: notificationpbgrpc.New(
			pb.NewNotificationPbClient(
				grpc_client.GetConnection(
					env.Env().GrpcClientUrl))),
		Notification_amqp:             notificationproducer.New(amqp_client),
		Json_place_holder_integration: json_place_holder.New(http_service),
		User_repository:               usersql.New(database),
		Users_cache:                   usersjsonplaceholdercache.New(cache_client),
	}
}