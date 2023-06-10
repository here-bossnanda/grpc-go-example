package main

import (
	"context"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, input *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet function was invoked with: %v\n", input)

	return &pb.SumResponse{
		Result: input.FirstNumber + input.SecondNumber,
	}, nil
}
