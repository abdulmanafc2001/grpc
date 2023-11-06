package main

import (
	"context"
	"log"

	pb "github.com/abdulmanafc2001/grpc/proto"
)

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("got request: %v", req.Name)
	return &pb.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}
