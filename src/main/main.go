package app

import (
	"log"
)

func Start() {
	err := New()

	if err != nil {
		log.Fatal("failed to create a server", err)
		return
	}
}
