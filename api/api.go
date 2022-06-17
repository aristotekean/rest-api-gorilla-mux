package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type API struct {
}

// schema convert the params to the data type
type BooksParams struct {
	Offset int `schema:"offset`
	Limit  int `schema:"limit`
}

type PostBook struct {
	Title string `json:"offset`
}

var (
	decoder = schema.NewDecoder()
	books   = []string{"Book 1", "Book 2", "Book 3"}
)

func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {

	params := &BooksParams{}

	err := decoder.Decode(params, r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Offset > len(books) || params.Offset < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Limit < 0 || params.Limit > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var from, to int
	if params.Offset > 0 {
		from = params.Offset
	}

	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(books)
	}

	json.NewEncoder(w).Encode(books[from:to])
}

func (a *API) getBook(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	idParam := pathParams["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idx := id - 1

	if idx < 0 || idx > len(books)-1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(books[idx])

}

func (a *API) postBook(w http.ResponseWriter, r *http.Request) {
	book := &PostBook{}

	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	books = append(books, book.Title)
	w.WriteHeader(http.StatusCreated)
}
