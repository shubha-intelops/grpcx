package main

import (
	log "github.com/sirupsen/logrus"
	"net"

	pb "github.com/shubha-intelops/grpcx/profile/gen/api/v1"
	grpccontrollers "github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/controllers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/client"
)

var (
	host = "localhost"
	port = "8081"
)

func main() {

	// Set up the TCP listener
	addr := net.JoinHostPort(host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Errorf("error while starting TCP listener: %s", err)
		return
	}

	log.Printf("TCP listener started at port: %s", port)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Create the Name server
	nameServer, err := grpccontrollers.NewNameServer()
	if err != nil {
		log.Errorf("error while creating nameServer: %s", err)
		return
	}
	// Register the Name server with the gRPC server
	pb.RegisterNameServiceServer(grpcServer, nameServer)

	// Enable reflection for the gRPC server
	reflection.Register(grpcServer)

	// Start serving gRPC requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Errorf("error serving gRPC: %s", err)
		return
	}

	// call external service here if the HasGrpcClients value is true
	// (that means this repo is a client to external service)
	client.Execute()

}
