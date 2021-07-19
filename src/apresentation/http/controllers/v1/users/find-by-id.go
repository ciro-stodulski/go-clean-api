package v1_user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (createController *createController) findById(gin_context *gin.Context) {
	user, err := createController.container.UserService.GetUser(uuid.New())

	if err != nil {
		gin_context.Status(http.StatusInternalServerError)
	}

	gin_context.JSON(http.StatusOK, user)
}
