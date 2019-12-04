package main

import (
	"time"
	"math/rand"
	"google.golang.org/grpc/credentials"
	"io"
	"flag"
	"log"
	"context"
	"google.golang.org/grpc"
	pb "06.go-rpc/02.route_guide/routeguide"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
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

	client := pb.NewRouteGuideClient(conn)

	// GetFeature
	printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})

	// ListFeatures
	printFeatures(client, &pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
	})

	// RecordRoute
	runRecordRoute(client)

	// RouteChat
	runRouteChat(client)
}

func randomPoint(r *rand.Rand) *pb.Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &pb.Point{Latitude: lat, Longitude: long}
}

func printFeature(client pb.RouteGuideClient, point *pb.Point) {
	log.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	feature, err := client.GetFeature(context.Background(), point)
	if err != nil {
		log.Fatalf("%v.GetFeature(_)=_, %v", client, err)
	}
	log.Println(feature)
}

func printFeatures(client pb.RouteGuideClient, rect *pb.Rectangle) {
	log.Printf("Looking for features within %v", rect)
	stream, err := client.ListFeatures(context.Background(), rect);
	if err != nil {
		log.Fatalf("%v.ListFeatures(_)=_, %v", client, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_)=_, %v", client, err)
		}
		log.Println(feature)
	}
}

func runRecordRoute(client pb.RouteGuideClient) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2
	var points []*pb.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}
	log.Printf("Traversing %d points.", len(points))

	stream, err := client.RecordRoute(context.Background())
	if err != nil {
		log.Fatalf("%v.RecordRoute(_)=_, %v", client, err)
	}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("%v.send(%v) = %v", stream, point, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route Summary: %v", reply)
}

func runRouteChat(client pb.RouteGuideClient) {
	log.Println("runRouteChat......")
	notes := []*pb.RouteNote{
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}

	stream, err := client.RouteChat(context.Background())
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	log.Println("runRouteChat......RouteChat")
	waitc := make(chan struct{})
	go func() {
		log.Println("runRouteChat......go-routine")
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				log.Println("runRouteChat....io.EOF")
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()
	log.Println("runRouteChat......notes=", len(notes))
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("Failed to send a note : %v", err)
		}
		log.Println("runRouteChat, send")
	}
	stream.CloseSend()
	<-waitc
	log.Println("runRouteChat......DONE")
}
