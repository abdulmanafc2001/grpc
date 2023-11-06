package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/abdulmanafc2001/grpc/proto"
)

func callBiDirectionalStreaming(cli pb.HelloServiceClient, names *pb.Names) {
	log.Printf("start bidirectional streaming")
	stream, err := cli.BiDirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send names: %s", err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res.Message)
		}
		close(waitc)
	}()

	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Println("Finishing bidirectional streaming")

}
