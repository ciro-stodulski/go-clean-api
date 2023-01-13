package container

import (
	list_users "go-clean-api/cmd/core/use-case/list-user"
	"go-clean-api/cmd/main/container/factories"
)

type (
	Container struct {
		ListUsersUseCase list_users.ListUsersUseCase
	}
)

func New() *Container {
	container_config := newContainerConfig()

	infra_context := factories.MakeInfraContext(
		container_config.Http_client,
	)

	user_service := factories.MakeServiceContext(infra_context).User_service

	return &Container{
		ListUsersUseCase: list_users.New(user_service),
	}
}
