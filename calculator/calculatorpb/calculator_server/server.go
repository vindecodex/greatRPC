package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vindecodex/gRPZ/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Add(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Add Invoked")
	firstNumber := req.GetInputs().GetFirstInput()
	secondNumber := req.GetInputs().GetSecondInput()
	result := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Server running")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
