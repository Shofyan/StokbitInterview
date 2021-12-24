package service

import (
	repository "github.com/Shofyan/StokbitInterview/repository/movie"
)

func (s *SvcMovie) GetOneMovie(imdbID string) repository.Movie {
	var m repository.Movie
	m = s.RepoMovie.GetOneMovie(imdbID)
	return m
}

func (s *SvcMovie) SearchMovie(key string, page string) repository.SeachMovie {
	var m repository.SeachMovie
	m = s.RepoMovie.SearchMovie(key, page)
	return m
}
