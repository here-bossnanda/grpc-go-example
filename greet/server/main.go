package main

import (
	"log"
	"net"

	pb "github.com/here-bossnanda/grpc-go-example/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:9001"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve : %v\n", err)
	}
}
