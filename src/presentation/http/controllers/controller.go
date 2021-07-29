package controllers

type Controller interface {
	LoadRoutes() []CreateRoute
	PathGroup() string
}

type CreateRoute struct {
	Method string
	Path   string
	Dto    interface{}
	Handle func(req HttpRequest) (HttpResponse, HttpResponseError)
}

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (ps Params) Get(name string) string {
	for _, entry := range ps {
		if entry.Key == name {
			return entry.Value
		}
	}
	return ""
}

type HttpRequest struct {
	Body    interface{}
	Params  Params
	Query   map[string][]string
	Headers map[string][]string
}

type Header struct {
	Key   string
	Value string
}

type HttpResponse struct {
	Data    interface{}
	Status  int
	Headers []Header
}

type HttpError struct {
	Code    string
	Message string
}
type HttpResponseError struct {
	Data   HttpError
	Status int
}
