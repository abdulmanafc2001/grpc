package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/abdulmanafc2001/grpc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial localhost:9000: %s", err)
	}
	defer conn.Close()
	cli := pb.NewHelloServiceClient(conn)

	// callSayHello(cli)
	names := &pb.Names{
		Name: []string{"Manaf","Alice","Bob"},
	}

	callSayHelloServerSteraming(cli,names)

}
