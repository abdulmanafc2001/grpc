package main

import (
	"context"
	"log"
	"time"

	pb "github.com/abdulmanafc2001/grpc/proto"
)

func callClientStreaming(cli pb.HelloServiceClient, names *pb.Names) {
	log.Println("client streaming started")
	stream, err := cli.ClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send")
	}
	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending: %v", err)
		}
		log.Printf("sent the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Println("Client streaming finished")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Message)
}
