package main

import (
	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest"
	"log"
	"net/http"
	"os"
)

func main() {

	port, isSet := os.LookupEnv("PORT")
	if !isSet {
		port = "8000" // default value, if not set at env level
	}
	router := rest.NewApiRestRouter()

	// TODO: repo init ...

	// Start the HTTP Server.
	log.Printf("Starting HTTP Server listening on :%s ...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
