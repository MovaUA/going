package main

import (
	"context"
	"fmt"
	"os"

	pb "github.com/movaua/going/going"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func newServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &greeterServer{})
	reflection.Register(server)
	return server
}

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("request: %v\n", req)
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %s from %s", req.Name, hostname)}, nil
}
