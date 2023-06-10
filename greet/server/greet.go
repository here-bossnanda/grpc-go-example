package main

import (
	"context"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/greet/proto"
)

func (s *Server) Greet(ctx context.Context, input *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with: %v\n", input)

	return &pb.GreetResponse{
		Result: "Hello " + input.FirstName,
	}, nil
}
