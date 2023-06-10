package main

import (
	"context"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/calculator/proto"
)

func doAvg(c pb.ServiceCalculatorClient) {
	log.Println("do Avg was invoked")

	numbers := []int32{3, 5, 9, 54, 23}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, number := range numbers {
		log.Printf("Sending req: %v\n", number)
		stream.Send(&pb.AVGRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}
