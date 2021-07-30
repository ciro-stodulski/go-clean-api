package middlewares

import (
	"go-api/src/presentation/http/controllers"
	"log"
)

func Log(req controllers.HttpRequest) {
	log.Default().Print("middleware started with succeffully")

	req.Next()
}
