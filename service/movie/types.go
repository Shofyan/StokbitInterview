package service

import (
	repository "github.com/Shofyan/StokbitInterview/repository/movie"
)

type SvcMovie struct {
	RepoMovie repository.RepoMovie
}

func New() *SvcMovie {
	return &SvcMovie{
		RepoMovie: *repository.New(),
	}
}
