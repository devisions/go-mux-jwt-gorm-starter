package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest"
	"github.com/devisions/go-mux-jwt-gorm-starter/users"
)

const (
	host     = "localhost"
	port     = 54325
	user     = "starter"
	password = "starter"
	dbname   = "go-mux-jwt-gorm-starter"
)

func main() {

	// Store init.
	dbConnInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	userSvc, err := users.NewUserService(dbConnInfo)
	if err != nil {
		log.Fatalf("Failed to init the user service: %v", err)
	}
	defer userSvc.Uninit()

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
