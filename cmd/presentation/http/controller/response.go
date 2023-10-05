package controller

type (
	HttpResponse[T any] struct {
		Data    T
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
		Detail  any    `json:"detail,omitempty"`
	}
)
