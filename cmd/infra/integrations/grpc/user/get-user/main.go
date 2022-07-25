package getuserservice

type getUserService struct {
	service PbGetUserService
}

func New(service PbGetUserService) GetUserService {

	return &getUserService{
		service: service,
	}
}
