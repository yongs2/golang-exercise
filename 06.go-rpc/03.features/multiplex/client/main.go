package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	hwpb "06.go-rpc/01.helloworld/helloworld"
	ecpb "06.go-rpc/03.features/proto/echo"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func callSayHello(c hwpb.GreeterClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &hwpb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("client.SayHello(_) = _, %v", err)
	}
	fmt.Println("Greeting: ", r.Message)
}

func calUnaryEcho(client ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("client.UnaryEcho(_) = _, %v:", err)
	}
	fmt.Println("UnaryEcho: ", resp.Message)
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	fmt.Println("--- calling helloworld.Greeter/SayHello ---")
	hwc := hwpb.NewGreeterClient(conn)
	callSayHello(hwc, "multiplex")

	fmt.Println()
	fmt.Println("--- calling routeguide.RouteGuide/GetFeature ---")

	rgc := ecpb.NewEchoClient(conn)
	calUnaryEcho(rgc, "this is a exampples/multiplex")
}
