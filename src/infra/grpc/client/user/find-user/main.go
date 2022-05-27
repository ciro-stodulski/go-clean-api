package find_user

import (
	"context"
	entity "go-api/src/core/entities"
	entity_user "go-api/src/core/entities/user"
	"go-api/src/infra/grpc/client/user/find-user/pb"
	"time"

	"google.golang.org/grpc"
)

type FindUserService struct {
	hostUrl    string
	connection *grpc.ClientConn
}

func New(connection *grpc.ClientConn) *FindUserService {
	return &FindUserService{
		hostUrl:    "localhost:50055",
		connection: connection,
	}
}

func (findUser *FindUserService) FindUser() (*entity_user.User, error) {
	client := pb.NewFindUserServiceClient(findUser.connection)

	request := &pb.NewRequestFindUser{
		ID: "1",
	}

	res, err := client.FindUser(context.Background(), request)

	if err != nil {
		return nil, err
	}

	data, err_data := time.Parse("2006-01-02", res.User.CreatedAt)

	if err_data != nil {
		return nil, err_data
	}

	return &entity_user.User{ID: entity.ConvertId(res.User.ID), Name: res.User.Name, Email: res.User.Email, CreatedAt: data}, nil
}
