package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/ninckblokje/golang-bits-n-bytes/RESTService/beer"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Syntappd hold the components to run this REST service
type Syntappd struct {
	l *log.Logger
	d beer.BeerStore
}

// New returns a new Gorilla ServeMux
// db should be converted to an interface / real database implementation
func New(beerstore beer.BeerStore) *Syntappd {
	// Setup Syntappd object
	app := Syntappd{
		l: log.New(os.Stdout, "", 0),
		d: beerstore,
	}

	return &app
}

// Healthcheck returns the health of the overall system
func (app *Syntappd) Healthcheck(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Status string
	}{
		Status: "OK",
	}

	w.Header().Set("Content-Type", "application/json")

	db, err := sqlx.Open("sqlite3", os.Getenv("PWD")+"/db.sqlite")
	if err = db.Ping(); err != nil {
		resp.Status = "NOK"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Home returns the homepage which is not really used
func (app *Syntappd) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Hello SynTouch`))
}

func (app *Syntappd) PostBeer(w http.ResponseWriter, r *http.Request) {
	var nb beer.NewBeer

	w.Header().Set("Content-Type", "application/json")

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Unable to read body" }`))
		return
	}

	err = json.Unmarshal(req, &nb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{ "error": "invalid request body" }`))
		return
	}

	rb, err := app.d.PostBeer(&nb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{ "error": "invalid input json" }`))
		return
	}
	json.NewEncoder(w).Encode(rb)

}

func (app *Syntappd) GetBeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	b, err := app.d.GetBeer(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "error": "beer does not exist" }`))
		return
	}

	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Unable to retrieve stored beer" }`))
		return
	}
}

func (app *Syntappd) GetAllBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	beer, err := app.d.GetAllBeers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Unable to retrieve stored beers" }`))
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(beer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Unable to retrieve stored beers" }`))
		return
	}
}
