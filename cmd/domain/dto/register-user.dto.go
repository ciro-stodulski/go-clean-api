package dto

type (
	RegisterUser struct {
		Name     string `json:"name"  binding:"required,min=3,max=10"`
		Email    string `json:"email"  binding:"required,min=0,max=50"`
		Password string `json:"password"  binding:"required,min=8,max=15"`
	}
)
