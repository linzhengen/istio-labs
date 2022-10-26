package main

import (
	"context"
	"flag"
	"fmt"
	"helloworld/protobuf/go/helloworld/v1"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.SayHelloRequest) (*helloworld.SayHelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.SayHelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
