package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ninckblokje/golang-bits-n-bytes/RESTService/handlers"
	"github.com/ninckblokje/golang-bits-n-bytes/RESTService/persistence"
)

func main() {
	db, err := persistence.Init()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}

	r := handlers.New(db)
	srv := http.Server{
		Handler: r,
		Addr:    "localhost:8080",
	}

	log.Fatal(srv.ListenAndServe())
}
