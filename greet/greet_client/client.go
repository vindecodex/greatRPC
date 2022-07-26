package main

import (
	"context"
	"fmt"
	"io"
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

	// res, err := unaryGreet(c)
	res, err := serverStreamGreet(c)
	if err != nil {
		log.Fatalf("Error invoking Greet RPC: %v", err)
	}

	log.Printf("Response value: %v", res)
}

func unaryGreet(c greetpb.GreetServiceClient) (*greetpb.GreetResponse, error) {
	fmt.Println("Invoke unary greet")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vincent",
			LastName:  "Villaluna",
		},
	}

	res, err := c.Greet(context.Background(), req)
	return res, err
}

func serverStreamGreet(c greetpb.GreetServiceClient) (greetpb.GreetService_GreetManyTimesClient, error) {
	fmt.Println("Invoke server streaming greet")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vincent",
			LastName:  "Villaluna",
		},
	}

	res, err := c.GreetManyTimes(context.Background(), req)
	for {
		result, err := res.Recv()
		if err == io.EOF {
			// end of stream
			break
		}
		if err != nil {
			log.Fatalf("Something is wrong: %v", err)
		}
		log.Printf("Response value: %v", result)
	}
	return res, err

}
