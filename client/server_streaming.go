package main

import (
	"context"
	"io"
	"log"

	pb "github.com/abdulmanafc2001/grpc/proto"
)

func callSayHelloServerSteraming(cli pb.HelloServiceClient, names *pb.Names) {
	log.Printf("%s", "stream started")
	stream, err := cli.ServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("%s", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%s", err)
		}
		log.Println(message)
	}
	log.Printf("%s", "streaming finished")
}
