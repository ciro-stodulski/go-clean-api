package mocks

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	user "go-clean-api/cmd/domain/entities/user"
)

func NewMockUser() *user.User {
	user, _ := user.New("test", "test", "test")
	return user
}

func CreateMockUser(name string, email string, password string) *user.User {
	user, _ := user.New(name, email, password)
	return user
}

func NewMockUserIntegration() []response_jsonplaceholder.User {
	return []response_jsonplaceholder.User{{
		Id:       12,
		Name:     "test",
		Username: "test",
		Email:    "test@test",
		Phone:    "test",
		Website:  "test",
		Address: response_jsonplaceholder.Address{
			Street:  "test",
			Suite:   "test",
			City:    "test",
			Zipcode: "test",
			Geo: response_jsonplaceholder.Geo{
				Lat: "test",
				Lng: "test",
			},
		},
		Company: response_jsonplaceholder.Company{
			Name:        "test",
			CatchPhrase: "test",
			Bs:          "test",
		},
	}}
}
