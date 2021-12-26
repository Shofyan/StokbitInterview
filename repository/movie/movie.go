package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (r *RepoMovie) GetOneMovie(imdbID string) (Movie, error) {
	// https://www.omdbapi.com/?apikey=faf7e5bb&i=tt4853102

	fmt.Println(imdbID)

	var m Movie
	resp, err := http.Get("https://www.omdbapi.com/?apikey=faf7e5bb&i=" + imdbID)
	if err != nil {
		log.Println(err)
		return m, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return m, err
	}
	json.Unmarshal(body, &m)

	return m, nil
}

func (r *RepoMovie) SearchMovie(key string, page string) (SeachMovie, error) {
	// https://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2
	var m SeachMovie

	fmt.Println(key)
	fmt.Println(page)

	resp, err := http.Get("https://www.omdbapi.com/?apikey=faf7e5bb&s=" + key + "&page=" + page)
	if err != nil {
		log.Println(err)
		return m, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return m, nil
	}

	json.Unmarshal(body, &m)
	return m, nil
}
