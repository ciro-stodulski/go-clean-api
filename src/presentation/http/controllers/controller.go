package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Register(gr *gin.RouterGroup)
}

func HealthCheck(context *gin.Context) {
	context.Status(http.StatusAccepted)
}
