package main

import (
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/calculator/proto"
)

func (s *Server) Prime(input *pb.PrimeRequest, stream pb.ServiceCalculator_PrimeServer) error {
	log.Printf("Prime function was invoked with: %v\n", input)

	number := input.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(
				&pb.PrimeResponse{
					Result: divisor,
				},
			)
			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
