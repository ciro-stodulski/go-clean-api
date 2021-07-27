package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Register(gr *gin.RouterGroup)
}
