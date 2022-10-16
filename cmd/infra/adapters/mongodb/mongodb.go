package adaptermongodb

import (
	"context"
	"go-clean-api/cmd/shared/env"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Database {
	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(env.Env().DBNoSqlHost)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx_timeout, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	err = client.Ping(ctx_timeout, nil)

	if err != nil {
		log.Fatal(err)
	}

	return client.Database("go-clean-api")
}
