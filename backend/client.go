package main

import (
	"context"
	"backend/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c := pb.NewHelloServiceClient(conn)

	response, err := c.Hello(context.Background(), &pb.HelloRequest{RequestText: "りくえすとおおおおおおおおおおおお"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Print(response.ResponseText)

	defer conn.Close()
}
