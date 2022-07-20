package find_service

import (
	"context"
	"go-api/cmd/interface/grpc/services/user/pb"
)

func (find_user_service *FindUserService) FindUser(ctx context.Context, req *pb.NewRequestFindUser) (*pb.NewResponseFindUser, error) {

	user, err := find_user_service.container.GetUserUseCase.GetUser(req.ID)

	if err != nil {
		return nil, err
	}

	return &pb.NewResponseFindUser{
		User: &pb.User{ID: user.ID.String(), Email: user.Email, Name: user.Name, CreatedAt: user.CreatedAt.String()},
	}, nil
}
