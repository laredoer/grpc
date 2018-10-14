package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/wule61/grpc/grpcusage"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "wule61"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewHelloClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	request, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(request.Message)
}
