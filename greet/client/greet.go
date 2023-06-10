package main

import (
	"context"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do Greet was invoked")

	res, err := c.Greet(
		context.Background(),
		&pb.GreetRequest{
			FirstName: "Nanda",
		})

	if err != nil {
		log.Fatalf("Couldn't not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
