package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	xw "github.com/sai-bhargav/grpc-https-gateway/proto/client"
	"github.com/sai-bhargav/grpc-https-gateway/server"

	"github.com/sai-bhargav/grpc-https-gateway/something"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9091", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	something.DoSomething()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible

	addr := "0.0.0.0:8001"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	xw.RegisterClientServiceServer(s, server.NewBackend())

	// Serve gRPC Server
	// log.Info("Serving gRPC on https://", addr)
	go func() {
		// log.Fatal(s.Serve(lis))

		fmt.Println("Fatal error in starting server", s.Serve(lis))

	}()

	mux := runtime.NewServeMux()
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}
	err = xw.RegisterClientServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	// oa := getOpenAPIHandler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "12000"
	}
	gatewayAddr := "0.0.0.0:" + port
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mux.ServeHTTP(w, r)
			return
		}),
	}

	// gwServer.TLSConfig = &tls.Config{
	// 	Certificates: []tls.Certificate{insecure.Cert},
	// }
	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
