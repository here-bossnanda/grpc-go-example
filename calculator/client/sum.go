package main

import (
	"context"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/calculator/proto"
)

func doSum(c pb.ServiceCalculatorClient) {
	log.Println("do sum was invoked")

	res, err := c.Sum(
		context.Background(),
		&pb.SumRequest{
			FirstNumber:  1,
			SecondNumber: 5,
		})

	if err != nil {
		log.Fatalf("Couldn't not sum: %v\n", err)
	}

	log.Printf("result: %d\n", res.Result)
}
