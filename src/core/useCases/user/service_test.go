package user

import (
	"testing"
	"time"

	entity "go-api/src/core/entities"
)

func newFixtureUser() *entity.User {
	return &entity.User{
		ID:        entity.NewID(),
		Name:      "oloco",
		Email:     "oloco@oloco.com",
		Password:  "213216574894",
		CreatedAt: time.Now(),
	}
}

func Test_SearchAndFind(t *testing.T) {
}
