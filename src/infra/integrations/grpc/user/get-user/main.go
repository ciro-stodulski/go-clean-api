package get_user_service

import (
	"context"
	"fmt"
	entity "go-api/src/core/entities"
	entity_user "go-api/src/core/entities/user"
	ports "go-api/src/core/ports"
	"go-api/src/infra/integrations/grpc/user/get-user/pb"
	"time"

	"google.golang.org/grpc"
)

type GetUserService interface {
	GetUser(context.Context, *pb.NewRequestGetUser, ...grpc.CallOption) (*pb.NewResponseGetUser, error)
}

type getUserService struct {
	service GetUserService
}

func New(service GetUserService) ports.GetUserService {

	return &getUserService{
		service: service,
	}
}

func (findUser *getUserService) GetUser(id string) (*entity_user.User, error) {
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
