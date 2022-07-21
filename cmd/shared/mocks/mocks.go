package mocks

import (
	user "go-api/cmd/core/entities/user"
	response_jsonplaceholder "go-api/cmd/infra/integrations/http/jsonplaceholder/responses"
)

func NewMockUser() *user.User {
	user, _ := user.New("test", "test", "test")
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
