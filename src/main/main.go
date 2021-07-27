package app

import (
	"log"
)

func StartApp() {
	server, err := New()

	if err != nil {
		log.Fatal("failed to create a server", err)
		return
	}

	err = server.ConnectToDabase()

	if err != nil {
		log.Fatal("failed to create to the database", err)
		return
	}

	defer server.CloseDB()

	err = server.Setup().engine.Start()

	if err != nil {
		log.Fatal("failed to serve the api", err)
		return
	}
}
