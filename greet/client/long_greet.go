package main

import (
	"context"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("do LongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Nanda"},
		{FirstName: "Novi"},
		{FirstName: "Ardy"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
