package service

import (
	"log"

	repository "github.com/Shofyan/StokbitInterview/repository/movie"
)

func (s *SvcMovie) GetOneMovie(imdbID string) (repository.Movie, error) {

	m, err := s.RepoMovie.GetOneMovie(imdbID)
	if err != nil {
		log.Println(err)
		return m, err
	}
	return m, nil
}

func (s *SvcMovie) SearchMovie(key string, page string) (repository.SeachMovie, error) {

	m, err := s.RepoMovie.SearchMovie(key, page)
	if err != nil {
		log.Println(err)
		return m, err
	}

	return m, nil
}
