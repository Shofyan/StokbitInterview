package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	svc "github.com/Shofyan/StokbitInterview/service/movie"
	"github.com/gorilla/mux"
)

func (rest *RestMovie) HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome home")
}

func (rest *RestMovie) GetOneMovie(w http.ResponseWriter, r *http.Request) {
	imdbID := mux.Vars(r)["id"]
	m := svc.GetOneMovie(imdbID)
	json.NewEncoder(w).Encode(m)
}

func (rest *RestMovie) SearchMovie(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	page := mux.Vars(r)["page"]
	m := svc.SearchMovie(key, page)
	json.NewEncoder(w).Encode(m)
}
