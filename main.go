package main

import (
	"fmt"
	"log"
	"net"

	pb "https://github.com/rahulsingh-21/dfn-internal-grpc/mfs"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server listening on port 50051")

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMfsServiceServer(grpcServer, &mfsServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
