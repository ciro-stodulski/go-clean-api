package user_use_case

//Service  interface
type Service struct {
	RepositoryUser  Repository
	IntegrationUser Integration
}

//NewService create new use case
func NewService(repository Repository, integration Integration) UseCase {
	return &Service{
		RepositoryUser:  repository,
		IntegrationUser: integration,
	}
}
