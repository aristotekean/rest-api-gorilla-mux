package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	r.Use(requestIDhandler)
	
	public := r.NewRoute().Subrouter()
	private := r.NewRoute().Subrouter()
	
	private.Use(authMiddleware)

	public.HandleFunc("/books", a.getBooks).Methods(http.MethodGet)
	public.HandleFunc("/books/{id}", a.getBook).Methods(http.MethodGet)

	private.HandleFunc("/books", a.postBook).Methods(http.MethodPost)

}
