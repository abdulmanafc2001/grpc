package main

import (
	"log"
	"net"

	pb "github.com/abdulmanafc2001/grpc/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("got error: %s", err)
	}
	log.Printf("server listening on: %s", lis.Addr())
	gRPCServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(gRPCServer, &server{})

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed serve: %s", err)
	}
}
