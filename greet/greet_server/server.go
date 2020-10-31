package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/kokiebisu/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	firstname := req.Greeting.GetFirstName()
	response := "Hello " + firstname
	res := &greetpb.GreetResponse{
		Result: response,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstname := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		fmt.Println(firstname, " entered", i, " times")
		res := &greetpb.GreetManyTimesResponse{
			Result: firstname,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil

}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}