package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	hwpb "06.go-rpc/01.helloworld/helloworld"
	ecpb "06.go-rpc/03.features/proto/echo"
)

var port = flag.Int("port", 50051, "the port to serve on")

type hwServer struct {
	hwpb.UnimplementedGreeterServer
}

func (s *hwServer) SayHello(ctx context.Context, in *hwpb.HelloRequest) (*hwpb.HelloReply, error) {
	return &hwpb.HelloReply{Message: "Hello " + in.Name}, nil
}

type ecServer struct {
	ecpb.UnimplementedEchoServer
}

func (s *ecServer) UnaryEcho(ctx context.Context, req *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	return &ecpb.EchoResponse{Message: req.Message}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	hwpb.RegisterGreeterServer(s, &hwServer{})
	ecpb.RegisterEchoServer(s, &ecServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
