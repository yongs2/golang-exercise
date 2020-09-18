package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "06.go-rpc/04.timer_event/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	port               = flag.Int("port", 9002, "The server port")
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

	myAddress := fmt.Sprintf("%s:%d", "localhost", *port)
	rspCreate, err := SendTimerCreate(client, &pb.TimerCreateRequest{
		CallbackUri: myAddress,
		SetTime:     time.Now().Format(time.RFC3339),
		ExpireSec:   10,
		Data:        "data_data",
		RepeatCount: -1,
	})
	if err != nil {
		log.Fatalf("failed to SendTimerCreate:%v", err)
	}
	log.Printf("rspCreate=[%v]\n", rspCreate)

	time.Sleep(time.Duration(1) * time.Second)

	rspDelete, err := SendTimerDelete(client, &pb.TimerDeleteRequest{
		TimerId: rspCreate.TimerId,
	})
	if err != nil {
		log.Fatalf("failed to SendTimerDelete:%v", err)
	}
	log.Printf("rspDelete=[%v]\n", rspDelete)
}

func SendTimerCreate(client pb.EventTimerClient, in *pb.TimerCreateRequest) (*pb.TimerCreateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("SendTimerCreate=[%v]\n", in)
	rsp, err := client.Create(ctx, in)
	if err != nil {
		log.Fatalf("%v.Create(_)=_, %v", client, err)
	}
	return rsp, nil
}

func SendTimerDelete(client pb.EventTimerClient, in *pb.TimerDeleteRequest) (*pb.TimerDeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("SendTimerDelete=[%v]\n", in)
	rsp, err := client.Delete(ctx, in)
	if err != nil {
		log.Fatalf("%v.Delete(_)=_, %v", client, err)
	}
	return rsp, nil
}
