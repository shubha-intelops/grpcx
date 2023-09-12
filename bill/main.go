package main

import (
	log "github.com/sirupsen/logrus"
	"net"

	pb "github.com/shubha-intelops/grpcx/bill/gen/api/v1"
	grpccontrollers "github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/controllers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	host = "localhost"
	port = "8082"
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

	// Create the Bill server
	billServer, err := grpccontrollers.NewBillServer()
	if err != nil {
		log.Errorf("error while creating billServer: %s", err)
		return
	}
	// Register the Bill server with the gRPC server
	pb.RegisterBillServiceServer(grpcServer, billServer)

	// Enable reflection for the gRPC server
	reflection.Register(grpcServer)

	// Start serving gRPC requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Errorf("error serving gRPC: %s", err)
		return
	}

}
