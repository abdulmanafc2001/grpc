package main

import (
	"io"
	"log"

	pb "github.com/abdulmanafc2001/grpc/proto"
)

func (s *server) BiDirectionalStreaming(stream pb.HelloService_BiDirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %s",req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}
