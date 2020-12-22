package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/xans-me/grpc-unary-example/protobuff"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) AddProduct(ctx context.Context, req *protobuff.ProductRequest) (*protobuff.ProductResponse, error) {
	fmt.Printf("Product function was invoked with %v", req)
	product := req.GetProduct()

	res := &protobuff.ProductResponse{
		Result: product.Name,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello I am Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protobuff.RegisterProductServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
