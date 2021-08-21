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

	err = server.db.ConnectToDabase()

	if err != nil {
		log.Fatal("failed to create to the database", err)
		return
	}

	defer server.db.CloseDB()

	servers := server.Setup()

	servers.http.Start()

	if err != nil {
		log.Fatal("failed to serve the api", err)
		return
	}
}
