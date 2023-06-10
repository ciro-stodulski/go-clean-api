package controller

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
		Detail  interface{}
	}
	HttpResponseError struct {
		Data   HttpError
		Status int
	}
)
