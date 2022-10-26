package main

import (
	"context"
	"helloworld/protobuf/go/helloworld/v1"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	defaultRpcEndpoint = "localhost:8080"
	rpcEndpoint        = os.Getenv("RPC_ENDPOINT")
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)
	if rpcEndpoint == "" {
		rpcEndpoint = defaultRpcEndpoint
	}
	conn, err := grpc.DialContext(
		context.Background(),
		rpcEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	gwmux := runtime.NewServeMux()
	if err := helloworld.RegisterGreeterServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}
	gwServer := &http.Server{
		Addr: ":8888",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				gwmux.ServeHTTP(w, r)
				return
			}
		}),
	}
	log.Fatal(gwServer.ListenAndServe())
}
