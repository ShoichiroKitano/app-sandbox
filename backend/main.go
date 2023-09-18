package main

import (
	"os/signal"
	"fmt"
	"net"
	"log"
	"context"
	"os"
	"google.golang.org/grpc"
	"backend/pb"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	helloService := &HelloService{}
	pb.RegisterHelloServiceServer(s, helloService)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

type HelloService struct {
}

func (s *HelloService) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println(request.RequestText)
	return &pb.HelloResponse {
		ResponseText: "aho",
	}, nil
}
