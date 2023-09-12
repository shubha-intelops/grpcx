package client

import (
	"context"
	pb "github.com/shubha-intelops/grpcx/profile/gen/api/v1"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Execute() {
	url := "localhost:8081"

	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Errorf("error occurred: %v", err)
		return
	}

	nameClient := pb.NewNameServiceClient(conn)

	ping, err := nameClient.Ping(context.Background(), &pb.NameRequest{})

	if err != nil {
		log.Errorf("error occurred: %v", err)
		return
	}

	log.Println(ping)
}
