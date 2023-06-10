package main

import (
	"context"
	"io"
	"log"

	pb "github.com/here-bossnanda/grpc-go-example/calculator/proto"
)

func doPrimes(c pb.ServiceCalculatorClient) {
	log.Println("do Primes was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Primes: %s\n", msg)
	}
}
