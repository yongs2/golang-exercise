package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "06.go-rpc/05.timer_event2/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:9001", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			//*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial:%v", err)
	}
	defer conn.Close()

	client := pb.NewEventTimerClient(conn)

	runTimerEvent(client)
}

func runTimerEvent(client pb.EventTimerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.TimerEvent(ctx)
	if err != nil {
		log.Fatalf("%v.TimerEvent(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				log.Printf("read EOF : %v\n", err)
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("Failed to receive a msg : %v\n", err)
				close(waitc)
				return
			}

			log.Printf("Got message %s\n", in)
			continue
		}
	}()

	createReq := &pb.TimerMsg{
		Command: &pb.TimerMsg_CreateReq{
			CreateReq: &pb.TimerCreateRequest{
				CallbackUri: "callback-0001",
				SetTime:     time.Now().Format(time.RFC3339),
				ExpireSec:   10,
				Data:        "datadata",
				RepeatCount: -1,
			},
		},
	}

	if err := stream.Send(createReq); err != nil {
		log.Fatalf("Failed to send a timerMsg: %v", err)
	}

	time.Sleep(time.Duration(10) * time.Second)

	log.Printf("client.CloseSend().....\n")
	stream.CloseSend()
	<-waitc
}
