package main

import (
	pb "github.com/wule61/grpc/grpcusage"

	"golang.org/x/net/context"

	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "hello " + in.Name,
	}, nil
}

func main() {
	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("grpc server listening at: 50051 port")
	server := grpc.NewServer()
	pb.RegisterHelloServer(server, &Server{})
	server.Serve(conn)
}
