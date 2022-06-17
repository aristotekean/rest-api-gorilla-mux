package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aristotekean/gorilla-api/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// create api routes
	a := &api.API{}

	// register the routes
	a.RegisterRoutes(r)
	r.HandleFunc("/", handleIndex).Methods(http.MethodGet)

	srv := &http.Server{
		Addr: ":8081",
		Handler: r,
	}

	log.Println("Listening...")
	srv.ListenAndServe()

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{\"message\": \"hello world\"}")
}
