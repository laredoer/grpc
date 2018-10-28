package main

import (
	"runtime/debug"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	pb "github.com/wule61/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

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
	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			loggingInterceptor,
			RecoveryInterceptor,
		),
	}

	fmt.Println("grpc server listening at: 50051 port")
	server := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(server, &SearchService{})
	server.Serve(conn)
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()
	return handler(ctx, req)
}
