package evtclient

import (
	"context"
	"io"
	"log"
	"time"

	pb "06.go-rpc/05.timer_event2/api/proto"

	// "golang.org/x/xerrors"
	"google.golang.org/grpc"
)

type EvTimerClient struct {
	Conn   *grpc.ClientConn
	Cli    pb.EventTimerClient
	stream pb.EventTimer_TimerEventClient

	onConnectFn          func()
	onCreateResponseFn   func(*pb.TimerCreateResponse)
	onDeleteResponseFn   func(*pb.TimerDeleteResponse)
	onTimerEventReportFn func(*pb.TimerEventReport)
}

func NewEvTimerClient() (*EvTimerClient, error) {
	return &EvTimerClient{}, nil
}

func (c *EvTimerClient) SetConnectCallback(cb func()) {
	c.onConnectFn = cb
}

func (c *EvTimerClient) SetCreateResponseCallback(cb func(*pb.TimerCreateResponse)) {
	c.onCreateResponseFn = cb
}

func (c *EvTimerClient) SetDeleteResponseCallback(cb func(*pb.TimerDeleteResponse)) {
	c.onDeleteResponseFn = cb
}

func (c *EvTimerClient) SetTimerEventReportCallback(cb func(*pb.TimerEventReport)) {
	c.onTimerEventReportFn = cb
}

func (c *EvTimerClient) Dial(endpoint string, dialTimeout time.Duration, opts []grpc.DialOption) error {
	var dialCtx context.Context
	if dialTimeout != 0 {
		var cancelFn func()
		dialCtx, cancelFn = context.WithTimeout(context.Background(), dialTimeout)
		defer cancelFn()
	}

	log.Printf("2>> dial[%v].options[%v], dialCtx=[%v]\n", endpoint, len(opts), dialCtx)
	// conn, err := grpc.DialContext(dialCtx, endpoint, opts...)
	// if err != nil {
	// 	return xerrors.Errorf("unable to dial endpoint[%v]: %w", endpoint, err)
	// }
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		log.Fatalf("failed to dial:%v", err)
	}

	c.Conn = conn
	c.Cli = pb.NewEventTimerClient(conn)
	return nil
}

func (c *EvTimerClient) Close() error {
	var err error
	if c.Conn != nil {
		err = c.Conn.Close()
		c.Conn = nil
	}
	return err
}

func (c *EvTimerClient) Run(ctx context.Context) (err error) {
	c.stream, err = c.Cli.TimerEvent(ctx)
	if err != nil {
		log.Printf("%v.TimerEvent(_) = _, %v", c.stream, err)
		return err
	}

	if c.onConnectFn != nil {
		c.onConnectFn()
	}

	for {
		in, err := c.stream.Recv()
		if err == io.EOF {
			log.Printf("read EOF : %v\n", err)
			// read done.
			// close(waitc)
			// return
			break
		}
		if err != nil {
			log.Printf("Failed to receive a msg : %v\n", err)
			// close(waitc)
			break
		}

		log.Printf("Got message [%v]\n", in.GetCommand())

		if msg := in.GetCreateRsp(); msg != nil {
			if c.onCreateResponseFn != nil {
				c.onCreateResponseFn(msg)
			}
		}

		if msg := in.GetDeleteRsp(); msg != nil {
			if c.onDeleteResponseFn != nil {
				c.onDeleteResponseFn(msg)
			}
		}

		if msg := in.GetEventRpt(); msg != nil {
			if c.onTimerEventReportFn != nil {
				c.onTimerEventReportFn(msg)
			}
		}

		continue
	}
	log.Printf("client.CloseSend().....\n")
	c.stream.CloseSend()
	return err
}

func (c *EvTimerClient) Create(msg *pb.TimerCreateRequest) (err error) {
	msgReq := &pb.TimerMsg{
		Command: &pb.TimerMsg_CreateReq{
			CreateReq: msg,
		},
	}

	if err = c.stream.Send(msgReq); err != nil {
		log.Printf("Create>> Failed to send a timerMsg: %v", err)
	}

	return err
}

func (c *EvTimerClient) Delete(msg *pb.TimerDeleteRequest) (err error) {
	msgReq := &pb.TimerMsg{
		Command: &pb.TimerMsg_DeleteReq{
			DeleteReq: msg,
		},
	}

	if err = c.stream.Send(msgReq); err != nil {
		log.Printf("Delete>> Failed to send a timerMsg: %v", err)
	}

	return err
}
