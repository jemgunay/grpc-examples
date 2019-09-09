package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/jemgunay/grpc-test/proto"
)

type server struct{}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterPublisherServer(grpcServer, &server{})

	l, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Printf("failed to listen: %s", err)
		return
	}

	if err = grpcServer.Serve(l); err != nil {
		fmt.Printf("server has shut down: %s", err)
	}
}

var data = []string{"cat", "dog", "tiger"}

// One request, one response.
func (s *server) UnaryExample(ctx context.Context, p *pb.Request) (*pb.Response, error) {
	log.Printf("a client requested topic %s", p.Topic)

	return &pb.Response{Data: fmt.Sprintf("%v", data)}, nil
}

// One request, multiple responses.
func (s *server) ResponseStreamExample(p *pb.Request, stream pb.Publisher_ResponseStreamExampleServer) error {
	log.Printf("a client requested topic %s", p.Topic)

	// stream responses to client
	for _, d := range data {
		if err := stream.Send(&pb.Response{Data: d}); err != nil {
			return err
		}
	}
	return nil
}

// Multiple requests, one response.
func (s *server) RequestStreamExample(stream pb.Publisher_RequestStreamExampleServer) error {
	// stream requests from client
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to retrieve topic: %s", err)
		}
		log.Printf("a client requested topic %s", req.Topic)
	}

	// respond with single response
	return stream.SendAndClose(&pb.Response{Data: "ok"})
}

// Multiple requests, multiple responses.
func (s *server) BidirectionalStreamExample(stream pb.Publisher_BidirectionalStreamExampleServer) error {
	// stream requests from client
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to retrieve topic: %s", err)
		}
		log.Printf("a client requested topic %s", req.Topic)
	}

	// stream responses to client
	for _, d := range data {
		if err := stream.Send(&pb.Response{Data: d}); err != nil {
			return err
		}
	}

	return nil
}
