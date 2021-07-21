package v1_user

import (
	entity_user "go-api/src/core/entities/user"

	entity "go-api/src/core/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (createController *createController) findById(gin_context *gin.Context) {
	id := entity.ConvertId(gin_context.Param("id"))

	user, err := createController.container.UserService.GetUser(id)

	if err != nil {
		if err == entity_user.ErrUserNotFound {
			gin_context.JSON(http.StatusBadRequest, gin.H{"code": "USER_NOT_FOUND", "message": entity_user.ErrUserNotFound.Error()})
			return
		}

		if err == entity_user.ErrIncorrectPassword {
			gin_context.JSON(http.StatusUnauthorized, gin.H{"code": "UNAUTHORIZED", "message": entity_user.ErrIncorrectPassword.Error()})
			return
		}

		gin_context.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_ERROR", "message": "internal error"})
		return
	}

	gin_context.JSON(http.StatusOK, user)
}
