package app

import (
	"log"
)

func Start() {
	err := New()

	if err != nil {
		log.Fatal("Error: Failed to start app", err)
		return
	}
}
