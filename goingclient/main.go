package main

import (
	pb "github.com/movaua/going/going"

	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

func main() {
	port, err := getPort()
	if err != nil {
		log.Fatalf("can't get port: %v\n", err)
	}
	log.Printf("using port %d\n", port)

	address := fmt.Sprintf("localhost:%d", port)

	log.Printf("connecting to %q\n", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
