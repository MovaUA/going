package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
)

func main() {
	port, err := getPort()
	if err != nil {
		log.Fatalf("can't get port: %v\n", err)
	}
	log.Printf("using port %d\n", port)

	address := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can't listen %q: %v\n", address, err)
	}
	log.Printf("listening %q\n", address)

	server := newServer()
	stopped := make(chan error)

	go func() {
		defer close(stopped)
		log.Println("serving server")
		stopped <- server.Serve(lis)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		<-stop
		log.Println("stopping server")
		server.Stop()
	}()

	if err := <-stopped; err != nil {
		log.Fatalf("can't serve server: %v\n", err)
	}

	log.Println("server is stopped")
}
