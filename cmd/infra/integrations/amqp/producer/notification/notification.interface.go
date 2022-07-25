package notificationproducer

type (
	Dto struct {
		Name  string `json:"name"`
		Event string `json:"event"`
	}

	NotificationProducer interface {
		SendNotify(dto Dto) error
	}
)
