package middlewares

import (
	"go-clean-api/cmd/presentation/http/controllers"
	"log"
)

func Global(req controllers.HttpRequest) {
	log.Default().Print("{Global} middleware call with succeffully")
	req.Headers["Add-Header-In-Middleware"] = []string{"Add in {Global} middleware and see in other middleware {Log}"}
	req.Next()
}
