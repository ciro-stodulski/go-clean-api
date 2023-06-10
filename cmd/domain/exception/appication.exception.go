package exception

type (
	ApplicationException struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func (appe *ApplicationException) Error() string {
	return appe.Message
}

func new(code string, message string) *ApplicationException {
	return &ApplicationException{code, message}
}
