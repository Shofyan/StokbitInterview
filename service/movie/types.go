package service

import (
	repository "github.com/Shofyan/StokbitInterview/repository/movie"
)

//go:generate mockgen -destination=movie_mock.go -package=service -source=types.go

type SvcMovie struct {
	RepoMovie repository.IRepoMovie
}

type ISvcMovie interface {
	GetOneMovie(imdbID string) (repository.Movie, error)
	SearchMovie(key string, page string) (repository.SeachMovie, error)
}

func New() *SvcMovie {
	return &SvcMovie{
		RepoMovie: repository.New(),
	}
}
