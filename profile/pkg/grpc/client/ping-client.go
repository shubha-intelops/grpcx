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

	profileDataClient := pb.NewProfileDataServiceClient(conn)

	ping, err := profileDataClient.Ping(context.Background(), &pb.ProfileDataRequest{})

	if err != nil {
		log.Errorf("error occurred: %v", err)
		return
	}

	log.Println(ping)
}
