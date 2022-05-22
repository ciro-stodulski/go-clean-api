package delete_user

import (
	"fmt"
	entity_root "go-api/src/core/entities"
)

func (service *deleteUserUseCase) DeleteUser(id string) error {
	id_uuid := entity_root.ConvertId(id)

	fmt.Println(id_uuid)
	err := service.RepositoryUser.DeleteById(id_uuid)

	if err != nil {
		return err
	}

	return nil
}
