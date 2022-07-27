package main

import (
	"context"
	pb "go-api/cmd/infra/integrations/grpc/notification/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (server *server) Verify(ctx context.Context, request *pb.Request) (*pb.Reponse, error) {

	return &pb.Reponse{
		Event: &pb.Event{
			Name:     "GRPC CLIENT connection with succeffully",
			Describe: "Test complete",
		},
	}, nil
}

func main() {
	serverStart()
}

func serverStart() {
	lis, err := net.Listen("tcp", "0.0.0.0:50044")

	if err != nil {
		log.Fatalf("error to list: %v", err)
	}

	grpc := grpc.NewServer()

	reflection.Register(grpc)

	pb.RegisterNotificationPbServer(grpc, &server{})
	log.Default().Println("GPRC SERVER CLIENT: started")

	if err := grpc.Serve(lis); err == nil {
		log.Fatalf("error to server: %v", err)
	}
}
