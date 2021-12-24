package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Shofyan/StokbitInterview/handler/rest"
)

func main() {
	restInstance := rest.New()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", restInstance.HomeLink)
	router.HandleFunc("/movie/{id}", restInstance.GetOneMovie).Methods("GET")
	router.HandleFunc("/movie/{key}/{page}", restInstance.SearchMovie).Methods("GET")

	println("start webserver at localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", router))

}
