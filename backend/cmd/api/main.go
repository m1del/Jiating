package main

import (
	"backend/internal/auth"
	"backend/internal/server"
	"fmt"
)

func main() {

	// initialize the auth package
	auth.Init()

	// initialize the server
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
