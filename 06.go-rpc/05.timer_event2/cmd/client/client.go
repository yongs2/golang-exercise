package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "06.go-rpc/05.timer_event2/api/proto"
	"06.go-rpc/05.timer_event2/evtclient"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:9001", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

var evtTimerClient *evtclient.EvTimerClient

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

	var err error

	evtTimerClient, err = evtclient.NewEvTimerClient()
	if err != nil {
		log.Fatalf("failed to create:%v", err)
	}
	defer evtTimerClient.Close()

	log.Printf("1>> dial.options[%v]\n", len(opts))
	if err = evtTimerClient.Dial(*serverAddr, 10, opts); err != nil {
		log.Fatalf("failed to dial:%v", err)
	}

	evtTimerClient.SetCreateResponseCallback(OnCreateResponse)
	evtTimerClient.SetDeleteResponseCallback(OnDeleteResponse)
	evtTimerClient.SetTimerEventReportCallback(OnTimerEventReport)

	go evtTimerClient.Run(context.Background())
	time.Sleep(time.Duration(1) * time.Second)

	createReq := pb.TimerCreateRequest{
		CallbackUri: "callback-0001",
		SetTime:     time.Now().Format(time.RFC3339),
		ExpireSec:   10,
		Data:        "datadata",
		RepeatCount: -1,
	}
	err = evtTimerClient.Create(&createReq)
	log.Printf("evtTimerClient.Create(%v).Err[%v]\n", createReq.CallbackUri, err)

	time.Sleep(time.Duration(10) * time.Second)
}

func OnCreateResponse(msg *pb.TimerCreateResponse) {
	if msg == nil {
		log.Printf("CreateResponse> msg is nil\n")
		return
	}

	log.Printf("CreateResponse.TimerId=[%v]\n", msg.GetTimerId())

	deleteReq := pb.TimerDeleteRequest{
		TimerId: msg.GetTimerId(),
	}
	err := evtTimerClient.Delete(&deleteReq)
	log.Printf("evtTimerClient.Delete(%v).Err[%v]\n", deleteReq.TimerId, err)
}

func OnDeleteResponse(msg *pb.TimerDeleteResponse) {
	if msg == nil {
		log.Printf("DeleteResponse> msg is nil\n")
		return
	}

	log.Printf("DeleteResponse.Result=[%v]\n", msg.GetResult())
}

func OnTimerEventReport(msg *pb.TimerEventReport) {
	if msg == nil {
		log.Printf("TimerEventReport> msg is nil\n")
		return
	}

	log.Printf("TimerEventReport.TimerId=[%v]\n", msg.GetTimerId())
}
