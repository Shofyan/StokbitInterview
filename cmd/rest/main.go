package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Shofyan/StokbitInterview/handler/rest"
	"github.com/Shofyan/StokbitInterview/util"
)

func main() {
	// select
	// u.ID,
	// u.UserName,
	// p.Username as ParentUserName
	// from USER u
	// left join USER p on u.Parent == p.ID
	fmt.Println("1. Query")
	fmt.Println("select u.ID, u.UserName,p.Username as ParentUserName from USER u left join USER p on u.Parent == p.ID")

	fmt.Println("3. Refactor")
	test := util.FindFirstStringInBracket("interview (lima kali) yes")
	fmt.Println(test)

	fmt.Println("4. Logic Test")
	in := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	result := util.Anagram(in)
	for _, v := range result {
		fmt.Println(v)
	}

	// soal nomor 2 membuat micro service
	restInstance := rest.New()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", restInstance.HomeLink)
	router.HandleFunc("/movie/{id}", restInstance.GetOneMovie).Methods("GET")
	router.HandleFunc("/movie/{key}/{page}", restInstance.SearchMovie).Methods("GET")

	println("start webserver at localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", router))

}
