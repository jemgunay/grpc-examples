package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/jemgunay/grpc-test/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:6000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to connect to server: %s", err)
		return
	}
	defer conn.Close()

	client := pb.NewPublisherClient(conn)
	unaryExample(client, "animals")
	responseStreamExample(client, "animals")
	requestStreamExample(client, "animals")
	bidirectionalStreamExample(client, "animals")
}

// One request, one response.
func unaryExample(client pb.PublisherClient, topic string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create request to server and receive single response
	resp, err := client.UnaryExample(ctx, &pb.Request{Topic: topic})
	if err != nil {
		log.Printf("failed to retrieve %s data: %s", topic, err)
		return
	}

	log.Printf("[%s]: %s", topic, resp.Data)
}

// One request, multiple responses.
func responseStreamExample(client pb.PublisherClient, topic string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create request to server
	stream, err := client.ResponseStreamExample(ctx, &pb.Request{Topic: topic})
	if err != nil {
		log.Printf("failed to subscribe to %s: %s", topic, err)
		return
	}

	// stream multiple responses from server
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to retrieve %s data: %s", topic, err)
			return
		}
		log.Printf("[%s]: %s", topic, resp.Data)
	}
}

// Multiple requests, one response.
func requestStreamExample(client pb.PublisherClient, topic string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create request to server
	stream, err := client.RequestStreamExample(ctx)
	if err != nil {
		log.Printf("failed to subscribe to %s: %s", topic, err)
		return
	}

	// stream multiple requests to the server
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.Request{Topic: topic}); err != nil {
			log.Printf("failed to send %s data: %s", topic, err)
			return
		}
	}

	// get single response from server
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("failed to retrieve %s data: %s", topic, err)
	}

	log.Printf("[%s]: %s", topic, resp.Data)
}

// Multiple requests, multiple responses.
func bidirectionalStreamExample(client pb.PublisherClient, topic string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create request to server
	stream, err := client.BidirectionalStreamExample(ctx)
	if err != nil {
		log.Printf("failed to subscribe to %s: %s", topic, err)
		return
	}

	// stream multiple requests to the server
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.Request{Topic: topic}); err != nil {
			log.Printf("failed to send %s data: %s", topic, err)
			return
		}
	}

	// signal end of requests
	if err := stream.CloseSend(); err != nil {
		log.Printf("failed to send close signal: %s", err)
		return
	}

	// stream multiple responses from server
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to retrieve %s data: %s", topic, err)
			return
		}
		log.Printf("[%s]: %s", topic, resp.Data)
	}
}
