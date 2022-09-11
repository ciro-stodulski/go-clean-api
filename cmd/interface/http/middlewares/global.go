package middlewares

import (
	ports_http "go-clean-api/cmd/interface/http/ports"
	"log"
)

func Global(req ports_http.HttpRequest) {
	log.Default().Print("{Global} middleware call with succeffully")
	req.Headers["Add-Header-In-Middleware"] = []string{"Add in {Global} middleware and see in other middleware {Log}"}
	req.Next()
}
