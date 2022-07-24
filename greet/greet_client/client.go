package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vindecodex/gRPZ/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client running...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vincent",
			LastName:  "Villaluna",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error invoking Greet RPC: %v", err)
	}

	log.Printf("Response value: %v", res)
}
