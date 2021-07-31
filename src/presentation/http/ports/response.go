package ports_http

type (
	HttpResponse struct {
		Data    interface{}
		Status  int
		Headers []Header
	}

	Header struct {
		Key   string
		Value string
	}

	HttpError struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	HttpResponseError struct {
		Data   HttpError
		Status int
	}
)
