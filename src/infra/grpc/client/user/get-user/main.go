package get_user_service

import (
	"context"
	"fmt"
	entity "go-api/src/core/entities"
	entity_user "go-api/src/core/entities/user"
	ports "go-api/src/core/ports"
	"go-api/src/infra/grpc/client/user/get-user/pb"
	"time"

	"google.golang.org/grpc"
)

type GetUserService struct {
	connection *grpc.ClientConn
	service    pb.GetUserServiceClient
}

func New(connection *grpc.ClientConn) ports.GetUserService {
	return &GetUserService{
		connection: connection,
		service:    pb.NewGetUserServiceClient(connection),
	}
}

func (findUser *GetUserService) GetUser(id string) (*entity_user.User, error) {
	request := &pb.NewRequestGetUser{
		ID: id,
	}

	res, err := findUser.service.GetUser(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res.Customer.CreatedAt)

	data, _ := time.Parse("2006-01-02 15:00", res.Customer.CreatedAt)
	fmt.Println(data)

	return &entity_user.User{ID: entity.ConvertId(res.Customer.ID), Name: res.Customer.Name, Email: res.Customer.Email, CreatedAt: data}, nil
}
