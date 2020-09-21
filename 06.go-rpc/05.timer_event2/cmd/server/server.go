package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "06.go-rpc/05.timer_event2/api/proto"
	"06.go-rpc/05.timer_event2/evtserver"

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

	evtTimerServer, err := evtserver.NewEvTimerServer(opts)
	if err != nil {
		log.Fatalf("failed to NewEvTimerServer:%v", err)
	}

	evtTimerServer.SetConnectCallback(OnConnect)
	evtTimerServer.SetDisconnectCallback(OnDisconnect)
	evtTimerServer.SetCreateRequestCallback(OnCreateRequest)
	evtTimerServer.SetDeleteRequestCallback(OnDeleteRequest)

	if err = evtTimerServer.Listen(lis); err != nil {
		panic(err)
	}
}

func OnConnect(stream pb.EventTimer_TimerEventServer) {
	serverTransportStream := grpc.ServerTransportStreamFromContext(stream.Context())
	log.Printf("OnConnect.Stream[%v],Method[%v]\n", stream, serverTransportStream.Method())
}

func OnDisconnect(stream pb.EventTimer_TimerEventServer) {
	serverTransportStream := grpc.ServerTransportStreamFromContext(stream.Context())
	log.Printf("OnDisconnect.Stream[%v],Method[%v]\n", stream, serverTransportStream.Method())
}

func OnCreateRequest(stream pb.EventTimer_TimerEventServer, msg *pb.TimerCreateRequest) {
	if msg == nil {
		log.Printf("CreateRequest> msg is nil\n")
		return
	}

	log.Printf("CreateRequest.CallbackUri=[%v]\n", msg.GetCallbackUri())

	createRsp := &pb.TimerMsg{
		Command: &pb.TimerMsg_CreateRsp{
			CreateRsp: &pb.TimerCreateResponse{
				CallbackUri: msg.CallbackUri,
				SetTime:     time.Now().Format(time.RFC3339),
				ExpireSec:   msg.ExpireSec,
				Data:        msg.Data,
				RepeatCount: msg.RepeatCount,
				TimerId:     "TimerId-0001",
			},
		},
	}
	if err := stream.Send(createRsp); err != nil {
		log.Fatalf("Failed to send a timerMsg: %v", err)
	}
}

func OnDeleteRequest(stream pb.EventTimer_TimerEventServer, msg *pb.TimerDeleteRequest) {
	if msg == nil {
		log.Printf("DeleteRequest> msg is nil\n")
		return
	}

	log.Printf("DeleteRequest.TimerId=[%v]\n", msg.GetTimerId())

	deleteRsp := &pb.TimerMsg{
		Command: &pb.TimerMsg_DeleteRsp{
			DeleteRsp: &pb.TimerDeleteResponse{
				Result: "Success",
			},
		},
	}
	if err := stream.Send(deleteRsp); err != nil {
		log.Fatalf("Failed to send a timerMsg: %v", err)
	}
}
