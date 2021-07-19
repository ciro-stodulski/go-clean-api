package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Register(gr *gin.RouterGroup)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
