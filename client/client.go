package main

import (
	"context"
	"log"

	pb "github.com/wule61/grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{Requets: "grpc"})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatalf("resp:%s", resp.GetResponse())
}
