package portsservice

type (
	Dto struct {
		Name  string `json:"name"`
		Event string `json:"event"`
	}

	NotificationService interface {
		SendNotify(dto Dto) error
		CheckNotify(msg string) (string error)
	}
)
