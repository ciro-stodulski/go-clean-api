package middlewares

import (
	"go-clean-api/cmd/presentation/http/controller"
	"log"
)

func Global(req controller.HttpRequest) {
	log.Default().Print("{Global} middleware call with succeffully")
	req.Headers["Add-Header-In-Middleware"] = []string{"Add in {Global} middleware and see in other middleware {Log}"}
	req.Next()
}
