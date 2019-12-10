package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "06.go-rpc/03.features/proto/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	const msg = "compress1"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: msg}, grpc.UseCompressor(gzip.Name))
	fmt.Printf("UnaryEcho called returned %q, %v\n", res.GetMessage(), err)
	if err != nil {
		log.Fatalf("Message=%q, err=%v; want Message=%q, err=<nil>", res.GetMessage(), err, msg)
	}
}
