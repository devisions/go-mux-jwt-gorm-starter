package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest"
	"github.com/devisions/go-mux-jwt-gorm-starter/users"
)

// Application Settings
const (
	host      = "localhost"
	port      = 54325
	user      = "starter"
	password  = "starter"
	dbName    = "go-mux-jwt-gorm-starter-db"
	dbMigrate = true
)

func main() {

	// Store init.
	dbConnInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	userStore, err := users.NewUserStore(dbConnInfo)
	if err != nil {
		log.Fatalf("Failed to init the user store: %s\n", err)
		return
	}
	if dbMigrate {
		err := userStore.Migrate()
		if err != nil {
			log.Fatalf("Failed to migrate the user store database: %s\n", err)
			return
		}
	}
	defer userStore.Close()
	userSvc := users.NewUserService(userStore)

	port, isSet := os.LookupEnv("PORT")
	if !isSet {
		port = "8000" // default value, if not set at env level
	}

	router := rest.NewApiRestRouter(userSvc)

	// Start the HTTP Server.
	log.Printf("Starting HTTP Server listening on :%s ...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
