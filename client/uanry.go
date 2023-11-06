package main

import (
	"context"
	"log"

	pb "github.com/abdulmanafc2001/grpc/proto"
)

func callSayHello(c pb.HelloServiceClient) {
	res, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Name:    "Manaf",
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	log.Println(res.Message)
}
