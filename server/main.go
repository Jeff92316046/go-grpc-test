package main

import (
	"context"
	"log"
	"net"

	"github.com/Jeff92316046/go-grpc-test/pb"
	"google.golang.org/grpc"
)

type MessageService struct {
}

func (m MessageService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Server say hello to " + in.GetName()}, nil
}

const (
	port = ":50051"
)

func main() {
	// Create gRPC Server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("gRPC server is running.")
	pb.RegisterGreeterServer(s, &MessageService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
