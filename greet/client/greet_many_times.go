package main

import (
	"context"
	"io"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("do GreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Nanda",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Greeting: %s\n", msg)
	}
}
