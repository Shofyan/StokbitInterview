package service

import (
	repository "github.com/Shofyan/StokbitInterview/repository/movie"
)

func (s *SvcMovie) GetOneMovie(imdbID string) repository.Movie {
	return s.RepoMovie.GetOneMovie(imdbID)
}

func (s *SvcMovie) SearchMovie(key string, page string) repository.SeachMovie {

	return s.RepoMovie.SearchMovie(key, page)
}
