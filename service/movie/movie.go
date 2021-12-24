package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	repository "github.com/Shofyan/StokbitInterview/repository/movie"
)

func GetOneMovie(imdbID string) repository.Movie {
	// https://www.omdbapi.com/?apikey=faf7e5bb&i=tt4853102
	var m repository.Movie
	resp, err := http.Get("https://www.omdbapi.com/?apikey=faf7e5bb&i=" + imdbID)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(body, &m)

	return m
}

func SearchMovie(key string, page string) repository.SeachMovie {
	// https://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2
	var m repository.SeachMovie

	fmt.Println(key)
	fmt.Println(page)

	resp, err := http.Get("https://www.omdbapi.com/?apikey=faf7e5bb&s=" + key + "&page=" + page)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &m)
	return m
}
