package grpc

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/Shofyan/StokbitInterview/handler/grpc/proto"
)

func TestServerNew(t *testing.T) {
	tests := []struct {
		name string
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ServerNew(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetOneMovie(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.GetOneMovieRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *pb.GetOneMovieResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetOneMovie(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetOneMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetOneMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_SearchMovie(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.SearchMovieRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *pb.SearchMovieResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SearchMovie(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.SearchMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.SearchMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}
