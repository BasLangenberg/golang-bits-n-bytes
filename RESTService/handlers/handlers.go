package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// New returns a new Gorilla ServeMux
func New(db *sqlx.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/health", healthcheck)

	return r
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	resp := struct{
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

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Hello SynTouch`))
}
