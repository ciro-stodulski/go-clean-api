package find_service

import (
	"context"
	"go-api/cmd/interface/grpc/services/user/pb"
	"go-api/cmd/main/container"
)

type findUserService struct {
	container *container.Container
}

func New(c *container.Container) *findUserService {
	return &findUserService{container: c}
}

func (find_user_service *findUserService) FindUser(ctx context.Context, req *pb.NewRequestFindUser) (*pb.NewResponseFindUser, error) {

	user, err := find_user_service.container.GetUserUseCase.GetUser(req.ID)

	if err != nil {
		return nil, err
	}

	return &pb.NewResponseFindUser{
		User: &pb.User{ID: user.ID.String(), Email: user.Email, Name: user.Name, CreatedAt: user.CreatedAt.String()},
	}, nil
}
