package middlewares

import (
	ports_http "go-api/src/interface/http/ports"
	"log"
)

func Log(req ports_http.HttpRequest) {
	log.Default().Print("Middleware {Log} in route call with succeffully")
	log.Default().Print(req.Headers["Add-Header-In-Middleware"][0])
	req.Next()
}
