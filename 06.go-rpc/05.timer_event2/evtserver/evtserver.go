package evtserver

import (
	"io"
	"log"
	"net"
	"time"

	pb "06.go-rpc/05.timer_event2/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type EvTimerServer struct {
	pb.UnimplementedEventTimerServer
	grpcServer *grpc.Server

	onConnectFn       func(stream pb.EventTimer_TimerEventServer)
	onDisconnectFn    func(stream pb.EventTimer_TimerEventServer)
	onCreateRequestFn func(stream pb.EventTimer_TimerEventServer, msg *pb.TimerCreateRequest)
	onDeleteRequestFn func(stream pb.EventTimer_TimerEventServer, msg *pb.TimerDeleteRequest)
}

func NewEvTimerServer(opts []grpc.ServerOption) (*EvTimerServer, error) {
	s := &EvTimerServer{}
	s.grpcServer = grpc.NewServer(opts...)
	pb.RegisterEventTimerServer(s.grpcServer, s)
	return s, nil
}

func (s *EvTimerServer) SetConnectCallback(cb func(stream pb.EventTimer_TimerEventServer)) {
	s.onConnectFn = cb
}

func (s *EvTimerServer) SetDisconnectCallback(cb func(stream pb.EventTimer_TimerEventServer)) {
	s.onDisconnectFn = cb
}

func (s *EvTimerServer) SetCreateRequestCallback(cb func(stream pb.EventTimer_TimerEventServer, msg *pb.TimerCreateRequest)) {
	s.onCreateRequestFn = cb
}

func (s *EvTimerServer) SetDeleteRequestCallback(cb func(stream pb.EventTimer_TimerEventServer, msg *pb.TimerDeleteRequest)) {
	s.onDeleteRequestFn = cb
}

func (s *EvTimerServer) Listen(lis net.Listener) (err error) {
	return s.grpcServer.Serve(lis)
}

func (s *EvTimerServer) TimerEvent(stream pb.EventTimer_TimerEventServer) error {
	log.Printf(">> Start Server.TimerEvent\n")

	if s.onConnectFn != nil {
		s.onConnectFn(stream)
	}

	ctx := stream.Context()
	for {
		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Read metadata from client.
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			log.Printf("BidirectionalStreamingEcho: failed to get metadata\n")
		}
		log.Printf("TimerEvent.metadata: [%v]\n", md)

		if t, ok := md["appid"]; ok {
			log.Printf("appid from metadata: [%v]\n", t)
		}

		// receive data from stream
		in, err := stream.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			break
		}
		if err != nil {
			log.Printf("receive error %v", err)
			break
		}
		log.Printf(">>> RecvMsg.Command[%v]\n", in.GetCommand())

		// Create and send header.
		header := metadata.New(map[string]string{"AppId": "TM2S", "timestamp": time.Now().Format(time.RFC3339)})
		stream.SendHeader(header)

		if msg := in.GetCreateReq(); msg != nil {
			if s.onCreateRequestFn != nil {
				s.onCreateRequestFn(stream, msg)
			}
		}

		if msg := in.GetDeleteReq(); msg != nil {
			if s.onDeleteRequestFn != nil {
				s.onDeleteRequestFn(stream, msg)
			}
		}
	}

	if s.onDisconnectFn != nil {
		s.onDisconnectFn(stream)
	}
	log.Printf("<< Stop Server.TimerEvent\n")
	return nil
}
