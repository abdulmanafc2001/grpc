package main

import (
	"log"
	"time"

	pb "github.com/abdulmanafc2001/grpc/proto"
)

func (s *server) ServerStreaming(req *pb.Names, stream pb.HelloService_ServerStreamingServer) error {
	log.Printf("request recieved: %s", req.Name)

	for _, name := range req.Name {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}
