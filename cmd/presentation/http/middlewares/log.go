package middlewares

import (
	"go-clean-api/cmd/presentation/http/controllers"
	"log"
)

func Log(req controllers.HttpRequest) {
	log.Default().Print("Middleware {Log} in route call with succeffully")
	log.Default().Print(req.Headers["Add-Header-In-Middleware"][0])
	req.Next()
}
