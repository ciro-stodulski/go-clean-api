package main

import (
	app "go-clean-api/cmd/main"
	"log"
)

func main() {
	log.Default().Print("Starting app")

	app.Start()
}
