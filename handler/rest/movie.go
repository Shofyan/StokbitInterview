package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (rest *RestMovie) HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome home")
}

func (rest *RestMovie) GetOneMovie(w http.ResponseWriter, r *http.Request) {
	imdbID := mux.Vars(r)["id"]
	m, err := rest.SvcMovie.GetOneMovie(imdbID)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}

func (rest *RestMovie) SearchMovie(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	page := mux.Vars(r)["page"]
	m, err := rest.SvcMovie.SearchMovie(key, page)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}
