package main

import (
	"context"
	"fmt"
	"log"

	"github.com/xans-me/grpc-unary-example/protobuff"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i am client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	defer conn.Close()

	c := protobuff.NewProductServiceClient(conn)
	// fmt.Printf("created client: %f",c)

	doUnary(c)

}

func doUnary(c protobuff.ProductServiceClient) {
	fmt.Println("Starting to do Unary RPC...")
	req := &protobuff.ProductRequest{
		Product: &protobuff.Product{
			Name:        "Kopi Kenangan Mantan",
			Description: "Kopi Espresso dengan Gula Aren",
			Price:       18000,
		},
	}

	resp, err := c.AddProduct(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Add product RPC: %v", err)
	}

	log.Printf("Response from Product Service: %s", resp.Result)
}
