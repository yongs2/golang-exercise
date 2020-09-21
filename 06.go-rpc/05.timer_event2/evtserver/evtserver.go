package evtserver

import (
	"io"
	"log"
	"net"

	pb "06.go-rpc/05.timer_event2/api/proto"

	"google.golang.org/grpc"
)

type EvTimerServer struct {
	pb.UnimplementedEventTimerServer
	grpcServer *grpc.Server

	onConnectFn       func(stream pb.EventTimer_TimerEventServer)
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

		// receive data from stream
		in, err := stream.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
		log.Printf(">>> RecvMsg.Command[%v]\n", in.GetCommand())

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
	log.Printf("<< Stop Server.TimerEvent\n")
	return nil
}
