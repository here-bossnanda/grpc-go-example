package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		fmt.Printf("Receiving: %v\n", req.FirstName)
		res += fmt.Sprintf("Hello %s\n", req.FirstName)
	}
}
