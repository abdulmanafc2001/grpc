package main

import (
	"io"
	"log"

	pb "github.com/abdulmanafc2001/grpc/proto"
)

func (s *server) ClientStreaming(stream pb.HelloService_ClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Messages{Message : messages})
		}
		if err != nil {
			return err
		}
		log.Printf("request: %v",req.Name)
		messages = append(messages,"Hello ",req.Name)
	}
}
