package middlewares

import (
	ports_http "go-api/src/presentation/http/controllers/ports"
	"log"
)

func Log(req ports_http.HttpRequest) {
	log.Default().Print("middleware call with succeffully")

	req.Next()
}
