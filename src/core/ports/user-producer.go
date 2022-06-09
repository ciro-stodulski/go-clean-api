package ports

type (
	CreateDto struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UserProducer interface {
		CreateUser(dto CreateDto) error
	}
)
