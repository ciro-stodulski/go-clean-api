package app

import (
	"fmt"
)

func StartApp() {
	server, err := New()

	if err != nil {
		fmt.Println("failed to create a server", err)
		return
	}

	err = server.Setup().Start()
	if err != nil {
		fmt.Println("failed to serve the api", err)
		return
	}
}
