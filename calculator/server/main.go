package main

import (
	"log"
	"net"

	pb "github.com/here-bossnanda/grpc-go-example/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:9001"

type Server struct {
	pb.ServiceCalculatorServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterServiceCalculatorServer(s, &Server{})
	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve : %v\n", err)
	}
}
