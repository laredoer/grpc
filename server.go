package main

import (
	pb "github.com/wule61/grpc/proto"
	"google.golang.org/grpc/credentials"

	"golang.org/x/net/context"

	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{
		Response: "hello " + in.Requets,
	}, nil
}

func main() {
	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	c, err := credentials.NewServerTLSFromFile("./conf/server.pem", "./conf/server.key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("grpc server listening at: 50051 port")
	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterSearchServiceServer(server, &SearchService{})
	server.Serve(conn)
}
