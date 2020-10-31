package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/kokiebisu/grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	c := calculatorpb.NewCalculatorServiceClient(conn)

	// doUnary(c)
	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Sum{
			NumberA: 3,
			NumberB: 10,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("I got a response back", res)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 120,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	var number string
	for {
		res, err := stream.Recv()
		fmt.Println("received result", res.GetResult())
		number = res.GetResult()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println("I got the result back", number)
}