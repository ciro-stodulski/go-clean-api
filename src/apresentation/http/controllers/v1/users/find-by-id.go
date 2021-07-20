package v1_user

import (
	entity "go-api/src/core/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (createController *createController) findById(gin_context *gin.Context) {
	id := entity.ConvertId(gin_context.Param("id"))

	user, err := createController.container.UserService.GetUser(id)

	if err != nil {
		gin_context.Status(http.StatusInternalServerError)
	}

	gin_context.JSON(http.StatusOK, user)
}
