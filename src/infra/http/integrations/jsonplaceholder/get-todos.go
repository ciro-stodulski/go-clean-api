package jsonplaceholder

import http_service "go-api/src/infra/http"

type JsonplaceholderIntegration struct {
	Engine http_service.HttpService
	url    string
}

func (intergration *JsonplaceholderIntegration) GetTodos() {
	intergration.Engine.CreateEngine(intergration.url).Engine.Send()
}
