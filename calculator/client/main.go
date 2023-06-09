package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/here-bossnanda/grpc-go-example/calculator/proto"
)

var addr string = "0.0.0.0:9002"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewServiceCalculatorClient(conn)

	// doSum(c)
	// doPrimes(c)
	doAvg(c)
}
