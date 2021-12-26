package rest

import (
	"net/http"
	"testing"
)

func TestRestMovie_HomeLink(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		rest *RestMovie
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rest.HomeLink(tt.args.w, tt.args.r)
		})
	}
}

func TestRestMovie_GetOneMovie(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		rest *RestMovie
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rest.GetOneMovie(tt.args.w, tt.args.r)
		})
	}
}

func TestRestMovie_SearchMovie(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		rest *RestMovie
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rest.SearchMovie(tt.args.w, tt.args.r)
		})
	}
}
