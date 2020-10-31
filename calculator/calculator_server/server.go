package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/kokiebisu/grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	numberA := req.Sum.GetNumberA()
	numberB := req.Sum.GetNumberB()

	fmt.Println("numberA", numberA)
	fmt.Println("numberB", numberB)

	sum := numberA + numberB
	fmt.Println("sum", sum)
	fmt.Println("sum string", strconv.Itoa(int(sum)))
	res := &calculatorpb.SumResponse{
		Result: strconv.Itoa(int(sum)),
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Println("hello")
	number := req.GetNumber()
	var k int32 = 2
	for number > 1 {
		if number % k == 0 {
			fmt.Println("factor", number)     // this is a factor
			number = number / k    // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}
		res := &calculatorpb.PrimeNumberDecompositionResponse{
			Result: strconv.Itoa(int(number)),
		}
		err := stream.Send(res)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func main() {
	fmt.Println("Successfully started server")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}