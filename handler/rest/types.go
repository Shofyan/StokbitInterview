package rest

import (
	service "github.com/Shofyan/StokbitInterview/service/movie"
)

type RestMovie struct {
	SvcMovie service.SvcMovie
}

func New() *RestMovie {
	return &RestMovie{
		SvcMovie: *service.New(),
	}
}
