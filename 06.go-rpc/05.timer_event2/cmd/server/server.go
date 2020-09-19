package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "06.go-rpc/05.timer_event2/api/proto"

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

func (s *eventTimerServer) TimerEvent(svr pb.EventTimer_TimerEventServer) error {
	log.Printf(">> Start Server.TimerEvent\n")

	ctx := svr.Context()
	for {
		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		msg, err := svr.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
		log.Printf(">>> RecvMsg=[%v], Command[%v]\n", msg, msg.GetCommand())
		if createReq := msg.GetCreateReq(); createReq != nil {
			createRsp := &pb.TimerMsg{
				Command: &pb.TimerMsg_CreateRsp{
					CreateRsp: &pb.TimerCreateResponse{
						CallbackUri: createReq.CallbackUri,
						SetTime:     time.Now().Format(time.RFC3339),
						ExpireSec:   createReq.ExpireSec,
						Data:        createReq.Data,
						RepeatCount: createReq.RepeatCount,
						TimerId:     "TimerId-0001",
					},
				},
			}
			if err := svr.Send(createRsp); err != nil {
				log.Fatalf("Failed to send a timerMsg: %v", err)
			}
		}
	}
	log.Printf("<< Stop Server.TimerEvent\n")
	return nil
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
