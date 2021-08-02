package user_use_case

//Service  interface
type Service struct {
	RepositoryUser Repository
}

//NewService create new use case
func NewService(repository Repository) UseCase {
	return &Service{
		RepositoryUser: repository,
	}
}
