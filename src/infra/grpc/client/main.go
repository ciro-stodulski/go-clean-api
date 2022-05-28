package grpc_client

import (
	"log"

	"google.golang.org/grpc"
)

func New(host string) *grpc.ClientConn {
	connection, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("error during connection ")
	}

	// melhorar estrutura para cruar uma instancia para o client
	// e poder tulixar o connection close
	//defer connection.Close()

	return connection
}
