package domainservice

type (
	Dto struct {
		Id    string `json:"id" bson:"_id,omitempty"`
		Name  string `json:"name"  bson:"name"`
		Event string `json:"event" bson:"event"`
	}

	NotificationService interface {
		SendNotify(dto Dto) error
		CheckNotify(msg string) (string error)
		SaveNotify(Dto) string
		FindById(id string) *Dto
	}
)
