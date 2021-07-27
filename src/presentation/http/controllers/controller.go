package controllers

import "github.com/gin-gonic/gin"

type CreateRoute struct {
	Method   string
	Path     string
	Function func(gin_context *gin.Context)
}

type Controller interface {
	LoadRoutes() []CreateRoute
	PathGroup() string
}
