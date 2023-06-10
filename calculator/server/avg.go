package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/calculator/proto"
)

func (s *Server) Avg(stream pb.ServiceCalculator_AvgServer) error {
	log.Println("Avg function was invoked")

	var sum int32 = 0
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AVGResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		sum += req.Number
		count++

		fmt.Printf("Receiving: %v\n", req.Number)
	}
}
