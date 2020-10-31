package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/kokiebisu/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm the client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	// doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Created Client", c)
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "ken",
			LastName: "okiebisu",
		},
	}
	res, err := c.Greet(context.Background(), req) 
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Response from greet", res)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("entered server streaming")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "ken",
		},
	}
	result, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}


	fmt.Println("sent")

	for {
		stream, err := result.Recv()
		if err == io.EOF {
			break;
		}
		if err != nil {
			log.Fatalln("something went wrong")
		}
		msg := stream.GetResult()
		fmt.Println("msg", msg)
	}

}