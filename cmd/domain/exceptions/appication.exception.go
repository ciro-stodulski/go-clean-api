package domainexceptions

type (
	applicationException struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func New(code string, message string) error {
	return &applicationException{code, message}
}

func (appe *applicationException) Error() string {
	var code = "code: " + appe.Code
	var message = "\nmessage: " + appe.Message

	return code + message
}
