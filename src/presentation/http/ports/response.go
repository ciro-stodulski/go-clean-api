package ports_http

type HttpResponse struct {
	Data    interface{}
	Status  int
	Headers []Header
}

type Header struct {
	Key   string
	Value string
}

type HttpError struct {
	Code    string
	Message string
}
type HttpResponseError struct {
	Data   HttpError
	Status int
}
