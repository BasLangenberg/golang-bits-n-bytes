package main

import (
	"github.com/gorilla/mux"
	"github.com/ninckblokje/golang-bits-n-bytes/RESTService/beer"
	"github.com/ninckblokje/golang-bits-n-bytes/RESTService/handlers"
	"log"
	"net/http"
)

func main() {
	// Will be used later in the tutorial
	//db, err := persistence.Init()
	//if err != nil {
	//	fmt.Printf("FATAL: %+v\n", err)
	//	os.Exit(1)
	//}

	app := handlers.New(beer.InMemoryBeerStore{})

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


	log.Fatal(srv.ListenAndServe())
}
