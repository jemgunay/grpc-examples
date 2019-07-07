package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/jemgunay/grpc-test/simple"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SimpleRPC(stream simple.SimpleService_SimpleRPCServer) error {
	log.Println("Started stream")
	for {
		in, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println("Got " + in.Msg)
	}
}

func main() {
	grpcServer := grpc.NewServer()
	simple.RegisterSimpleServiceServer(grpcServer, &server{})

	l, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Printf("failed to listen: %s", err)
		return
	}

	if err = grpcServer.Serve(l); err != nil {
		fmt.Printf("server has shut down: %s", err)
	}
}
