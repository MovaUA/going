package main

import (
	"crypto/tls"

	pb "github.com/movaua/going/going"

	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	defaultName = "world"
)

func main() {
	host, err := getHost()
	if err != nil {
		log.Fatalf("can't get host: %v\n", err)
	}
	log.Printf("using host %q\n", host)

	port, err := getPort()
	if err != nil {
		log.Fatalf("can't get port: %v\n", err)
	}
	log.Printf("using port %d\n", port)

	address := fmt.Sprintf("%s:%d", host, port)

	log.Printf("connecting to %q\n", address)
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)), grpc.WithTimeout(10*time.Second), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	log.Printf("connected to %q\n", address)
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}
	log.Printf("response: %v\n", r)
}
