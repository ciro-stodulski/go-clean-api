package factories

import (
	domainnotificationproducer "go-clean-api/cmd/domain/integrations/amqp"
	domainnotificationpbgrpc "go-clean-api/cmd/domain/integrations/grpc"
	domainjsonplaceholder "go-clean-api/cmd/domain/integrations/http"
	domainusersjsonplaceholdercache "go-clean-api/cmd/domain/repositories/cache"
	domainnotificationcollection "go-clean-api/cmd/domain/repositories/no-sql"
	domainusersql "go-clean-api/cmd/domain/repositories/sql"
	amqpclient "go-clean-api/cmd/infra/integrations/amqp"
	notificationproducer "go-clean-api/cmd/infra/integrations/amqp/notification"
	grpc_client "go-clean-api/cmd/infra/integrations/grpc"
	http_service "go-clean-api/cmd/infra/integrations/http"
	json_place_holder "go-clean-api/cmd/infra/integrations/http/jsonplaceholder"
	usersjsonplaceholdercache "go-clean-api/cmd/infra/repositories/cache/users-jsonplaceholder"
	notificationcollection "go-clean-api/cmd/infra/repositories/no-sql/notification"
	usersql "go-clean-api/cmd/infra/repositories/sql/user"

	notificationpbgrpc "go-clean-api/cmd/infra/integrations/grpc/notification"
	"go-clean-api/cmd/infra/integrations/grpc/notification/pb"
	cache_client "go-clean-api/cmd/infra/repositories/cache"
	"go-clean-api/cmd/shared/env"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	InfraContext struct {
		NotificationPbGrpc            domainnotificationpbgrpc.NotificationPbGrpc
		Notification_amqp             domainnotificationproducer.NotificationProducer
		Notification_collection       domainnotificationcollection.NotificationCollection
		Json_place_holder_integration domainjsonplaceholder.JsonPlaceholderIntegration
		User_repository               domainusersql.UserSql
		Users_cache                   domainusersjsonplaceholdercache.UsersJsonPlaceholderCache
	}
)

func MakeInfraContext(
	grpc_client grpc_client.GRPCClient,
	amqp_client amqpclient.AmqpClient,
	http_service http_service.HttpClient,
	database *gorm.DB,
	cache_client cache_client.CacheClient,
	no_sqldatabase *mongo.Database,
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
		Notification_collection:       notificationcollection.New(no_sqldatabase),
	}
}
