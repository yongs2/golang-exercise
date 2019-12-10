package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	ecpb "06.go-rpc/03.features/proto/echo"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/testdata"
)
// golang.org/x/oauth2 를 포함시켰는데도, undefined: metadata.OnGCE 라고 표시되면, 
// https://github.com/golang/oauth2 를 참조하여, cloud.google.com/go/compute/metadata 로 수정

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func callUnaryEcho(client ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("client.UnaryEcho(_)=_, err=%v", err)
	}
	fmt.Println("UnaryEcho:", resp.Message)
}

func main() {
	flag.Parse()

	perRPC := oauth.NewOauthAccess(fetchToken())
	creds, err := credentials.NewClientTLSFromFile(testdata.Path("ca.pem"), "x.test.youtube.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(creds),
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	rgc := ecpb.NewEchoClient(conn)

	callUnaryEcho(rgc, "Hello world")
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-toekn",
	}
}
