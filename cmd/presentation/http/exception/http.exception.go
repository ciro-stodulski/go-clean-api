package httpexception

type (
	httpError struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	HttpException struct {
		Data   httpError
		Status int
	}
)
