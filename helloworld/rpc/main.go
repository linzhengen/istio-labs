package main

import (
	"context"
	"flag"
	"fmt"
	"helloworld/protobuf/go/helloworld/v1"
	"log"
	"math/rand"
	"net"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.SayHelloRequest) (*helloworld.SayHelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	rand.Seed(time.Now().UnixNano())
	min := 6000
	max := 6700
	st := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(st) * time.Millisecond)
	return &helloworld.SayHelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	healthServer := health.NewServer()
	helloworld.RegisterGreeterServiceServer(s, &server{})
	healthpb.RegisterHealthServer(s, healthServer)
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully")
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		// close listeners to stop accepting new connections,
		// will block on any existing transports
		s.GracefulStop()
	}()
	select {
	case <-ch:
		log.Printf("Graceful stopped")
	case <-time.After(10 * time.Second):
		// took too long, manually close open transports
		// e.g. watch streams
		log.Printf("Graceful stop timeout, force stop!!")
		s.Stop()
		<-ch
	}
	log.Println("Server successfully stopped")
}
