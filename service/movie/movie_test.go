package service

import (
	"testing"

	repository "github.com/Shofyan/StokbitInterview/repository/movie"
	"github.com/stretchr/testify/assert"
)

func TestSvcMovie_GetOneMovie(t *testing.T) {
	type args struct {
		imdbID string
	}
	tests := []struct {
		name    string
		s       *SvcMovie
		args    args
		want    repository.Movie
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				imdbID: "tt4853102",
			},
			want: repository.Movie{
				Title: "Batman: The Killing Joke",
			},
			s: &SvcMovie{
				RepoMovie: repository.New(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetOneMovie(tt.args.imdbID)

			if tt.name == "success" {
				assert.Equal(t, got.Title, tt.want.Title)
			}

			if tt.wantErr {
				assert.NotNil(t, err)
			}

		})
	}
}
