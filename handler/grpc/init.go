package grpc

import (
	"context"
	"log"
	"strconv"

	pb "github.com/Shofyan/StokbitInterview/handler/grpc/proto"
	service "github.com/Shofyan/StokbitInterview/service/movie"
)

type Server struct {
	pb.UnimplementedStockbitServer
	SvcMovie service.SvcMovie
}

func ServerNew() *Server {
	return &Server{
		SvcMovie: *service.New(),
	}
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GetOneMovie(ctx context.Context, in *pb.GetOneMovieRequest) (*pb.GetOneMovieResponse, error) {
	log.Printf("Received: %v", in.GetId())

	imdbID := in.GetId()
	m, err := s.SvcMovie.GetOneMovie(imdbID)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	rating := []*pb.Rating{}
	for _, v := range m.Ratings {
		temp := pb.Rating{
			Source: v.Source,
			Value:  v.Value,
		}
		rating = append(rating, &temp)
	}

	return &pb.GetOneMovieResponse{
		Title:      m.Title,
		Year:       m.Year,
		Rated:      m.Year,
		Released:   m.Released,
		Runtime:    m.Runtime,
		Genre:      m.Genre,
		Director:   m.Director,
		Writer:     m.Writer,
		Actors:     m.Actors,
		Plot:       m.Plot,
		Language:   m.Language,
		Country:    m.Country,
		Awards:     m.Awards,
		Poster:     m.Poster,
		Ratings:    rating,
		Metascore:  m.Metascore,
		ImdbRating: m.ImdbRating,
		ImdbID:     m.ImdbRating,
		Type:       m.Type,
		DVD:        m.DVD,
		BoxOffice:  m.BoxOffice,
		Production: m.Production,
		Website:    m.Website,
		Response:   m.Response,
	}, nil
}

func (s *Server) SearchMovie(ctx context.Context, in *pb.SearchMovieRequest) (*pb.SearchMovieResponse, error) {

	key := in.GetKey()
	page := in.GetPage()

	m, err := s.SvcMovie.SearchMovie(key, page)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	movies := []*pb.MovieSummary{}
	for _, v := range m.Search {
		temp := &pb.MovieSummary{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		}
		movies = append(movies, temp)
	}

	res := &pb.SearchMovieResponse{
		Search:       movies,
		TotalResults: strconv.FormatInt(int64(len(movies)), 10),
		Response:     "True",
	}

	return res, nil
}
