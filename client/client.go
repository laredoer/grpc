package main

import (
	"context"
	"log"

	"google.golang.org/grpc/credentials"

	pb "github.com/wule61/grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	c, err := credentials.NewClientTLSFromFile("../conf/server.pem", "x")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{Requets: "grpc"})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Fatalf("resp:%s", resp.GetResponse())
}
