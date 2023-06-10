package main

import (
	"fmt"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/greet/proto"
)

func (s *Server) GreetManyTimes(input *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", input)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", input.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
