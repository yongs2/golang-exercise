package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "06.go-rpc/04.timer_event/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	port     = flag.Int("port", 9001, "The server port")
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
)

type eventTimerServer struct {
	pb.UnimplementedEventTimerServer
}

func (s *eventTimerServer) Create(ctx context.Context, req *pb.TimerCreateRequest) (*pb.TimerCreateResponse, error) {
	log.Printf(">> Server.CreateRequest=[%v]\n", req)
	rsp := pb.TimerCreateResponse{
		CallbackUri: req.CallbackUri,
		SetTime:     req.SetTime,
		ExpireSec:   req.ExpireSec,
		Data:        req.Data,
		RepeatCount: req.RepeatCount,
		TimerId:     "timerid-0001",
	}
	log.Printf("<< Server.CreateResponse=[%v]\n", rsp)
	return &rsp, nil
}

func (s *eventTimerServer) Delete(ctx context.Context, req *pb.TimerDeleteRequest) (*pb.TimerDeleteResponse, error) {
	log.Printf(">> Server.TimerDeleteRequest=[%v]\n", req)
	rsp := pb.TimerDeleteResponse{
		Result: "Success",
	}
	log.Printf("<< Server.TimerDeleteResponse=[%v]\n", rsp)
	return &rsp, nil
}

func newServer() *eventTimerServer {
	s := &eventTimerServer{}
	return s
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEventTimerServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
