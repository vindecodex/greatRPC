package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vindecodex/gRPZ/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client running...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	req := &calculatorpb.SumRequest{
		Inputs: &calculatorpb.Input{
			FirstInput:  10,
			SecondInput: 3,
		},
	}

	res, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Error invoking Greet RPC: %v", err)
	}

	log.Printf("Resoonse value: %v", res)
}
