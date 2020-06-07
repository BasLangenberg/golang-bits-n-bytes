package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ninckblokje/golang-bits-n-bytes/RESTService/handlers"
	"github.com/ninckblokje/golang-bits-n-bytes/RESTService/persistence"
	"log"
	"net/http"
	"os"
)

func main() {
	sqlite, err := persistence.Init()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}

	app := handlers.New(sqlite)

	// Setup Handlers
	r := mux.NewRouter()
	r.HandleFunc("/", app.Home)
	r.HandleFunc("/health", app.Healthcheck)
	r.HandleFunc("/beer", app.PostBeer).Methods("POST")
	r.HandleFunc("/beer", app.GetAllBeers).Methods("GET")
	r.HandleFunc("/beer/{id}", app.GetBeer).Methods("GET")


	srv := http.Server{
		Handler: r,
		Addr:    "localhost:8080",
	}

	// This could be improved by using a "global" logger hanging on the app struct
	log.Println("application now running on localhost:8080")

	log.Fatal(srv.ListenAndServe())
}
