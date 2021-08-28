package producer

type Body struct {
	Data interface{}
}

type Producer interface {
	Send(
		message Body,
	) error
}
