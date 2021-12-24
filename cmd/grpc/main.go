package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	stGRPC "github.com/Shofyan/StokbitInterview/handler/grpc"
	pb "github.com/Shofyan/StokbitInterview/handler/grpc/proto"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStockbitServer(s, &stGRPC.Server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
