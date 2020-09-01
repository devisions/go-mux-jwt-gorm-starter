## Starter Project on Go using Mux, JWT, and GORM

This is a starter project for a GoLang implementation of an HTTP (such as RESTful) API using:

- Gorilla Mux - as the HTTP router
- JWT - JSON Web Tokens, one of the popular standard for protecting endpoints
- GORM - the ORM layer, used for a high-level interaction with the database

### Setup

In terms of prerequisites:

- The dependencies are managed by Go Modules, so this is taken care implicitly (no manual action req'd)
- For restarting on changes (the process is restarted, we don't aim for "live reloading" now), [air](https://github.com/cosmtrek/air) utility is used in the run script. Just install this tool using `go get -u github.com/cosmtrek/air`.
