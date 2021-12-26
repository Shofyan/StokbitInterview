package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoMovie_SearchMovie(t *testing.T) {
	type args struct {
		key  string
		page string
	}
	tests := []struct {
		name string
		r    *RepoMovie
		args args
		want SeachMovie
		err  bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				key:  "batman",
				page: "2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.SearchMovie(tt.args.key, tt.args.page)

			if err == nil {
				assert.NotNil(t, got)
			} else {
				assert.NotNil(t, err)
			}

		})
	}
}

func TestRepoMovie_GetOneMovie(t *testing.T) {
	type args struct {
		imdbID string
	}
	tests := []struct {
		name string
		r    *RepoMovie
		args args
		want Movie
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				imdbID: "tt4853102",
			},
			want: Movie{
				Title: "Batman: The Killing Joke",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetOneMovie(tt.args.imdbID)
			if err == nil {
				assert.Equal(t, got.Title, tt.want.Title)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
