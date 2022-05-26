package find_service

import (
	"context"
	"go-api/src/presentation/grpc/services/user/find-user/pb"
)

func (find_user_service *FindUserService) FindUser(ctx context.Context, req *pb.NewRequestFindUser) (*pb.NewResponseFindUser, error) {

	user, _ := find_user_service.container.GetUserUseCase.GetUser(req.ID)

	return &pb.NewResponseFindUser{
		User: &pb.User{ID: user.ID.String(), Email: user.Email, Name: user.Name, CreatedAt: user.CreatedAt.String()},
	}, nil
}
