package app

import (
	"log"
)

func Start() {
	err := New()

	if err != nil {
		log.Fatal("Failed to start app", err)
		return
	}
}
