package main

import (
	"log"
	"net/http"
)

func main() {

	// handlers holds all the needed objects in the app
	handlers := new(Handlers)
	db := new(Database)
	// Start DB. Will execute migrations if needed
	handlers.db = db.ConnectDB()

	// Set routes and start the server in port 80 (docker will redirect)
	handlers.setRoutes()
	log.Fatal(http.ListenAndServe(":80", handlers.router))
}
