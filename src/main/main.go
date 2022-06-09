package app

import (
	"log"
)

func StartApp() {
	app, err := New()

	if err != nil {
		log.Fatal("failed to create a server", err)
		return
	}

	err = app.db.ConnectToDabase()

	if err != nil {
		log.Fatal("failed to create to the database", err)
		return
	}

	defer app.db.CloseDB()

	server := app.Setup()

	server.http.Start()

	if err != nil {
		log.Fatal("failed to serve the api", err)
		return
	}
}
