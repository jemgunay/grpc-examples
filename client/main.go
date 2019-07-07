package main

import (
	"fmt"

	pb "github.com/jemgunay/grpc-test/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:6000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to connect to server: %s", err)
		return
	}
	defer conn.Close()

	client := pb.NewSimpleServiceClient()
}
