package controllers

import (
	"context"

	pb "github.com/shubha-intelops/grpcx/profile/gen/api/v1"
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/models"
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/services"
)

type NameServer struct {
	nameService *services.NameService
	pb.UnimplementedNameServiceServer
}

func NewNameServer() (*NameServer, error) {
	nameService, err := services.NewNameService()
	if err != nil {
		return nil, err
	}
	return &NameServer{
		nameService: nameService,
	}, nil
}

func (*NameServer) Ping(_ context.Context, _ *pb.NameRequest) (*pb.NameResponse, error) {
	return &pb.NameResponse{
		Message: "Server is healthy and working!",
	}, nil
}

func (nameServer *NameServer) CreateName(_ context.Context, req *pb.CreateNameRequest) (*pb.CreateNameResponse, error) {
	name := models.Name{

		Address: req.Name.GetAddress(),

		Age: req.Name.GetAge(),

		Email_address: req.Name.GetEmail_address(),

		Name: req.Name.GetName(),
	}

	if _, err := nameServer.nameService.CreateName(&name); err != nil {
		return nil, err
	}

	return &pb.CreateNameResponse{
		Message: "Name Created Successfully!",
	}, nil
}

func (nameServer *NameServer) GetName(_ context.Context, req *pb.GetNameRequest) (*pb.GetNameResponse, error) {
	id := req.GetId()
	retrievedName, err := nameServer.nameService.GetName(id)

	if err != nil {
		return nil, err
	}

	nameResponse := &pb.Name{
		Id: id,

		Address: retrievedName.Address,

		Age: retrievedName.Age,

		Email_address: retrievedName.Email_address,

		Name: retrievedName.Name,
	}

	return &pb.GetNameResponse{
		Name: nameResponse,
	}, nil
}

func (nameServer *NameServer) ListNames(_ context.Context, _ *pb.ListNamesRequest) (*pb.ListNamesResponse, error) {
	names, err := nameServer.nameService.ListNames()

	if err != nil {
		return nil, err
	}

	var nameList []*pb.Name
	for _, retrievedName := range names {
		nameResponse := &pb.Name{
			Id: retrievedName.Id,

			Address: retrievedName.Address,

			Age: retrievedName.Age,

			Email_address: retrievedName.Email_address,

			Name: retrievedName.Name,
		}
		nameList = append(nameList, nameResponse)
	}

	return &pb.ListNamesResponse{
		Name: nameList,
	}, nil
}

func (nameServer *NameServer) UpdateName(_ context.Context, req *pb.UpdateNameRequest) (*pb.UpdateNameResponse, error) {
	id := req.GetId()

	nameRequest := models.Name{
		Id: id,

		Address: req.Name.GetAddress(),

		Age: req.Name.GetAge(),

		Email_address: req.Name.GetEmail_address(),

		Name: req.Name.GetName(),
	}
	_, err := nameServer.nameService.UpdateName(id, &nameRequest)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateNameResponse{
		Message: "Name Updated Successfully!",
	}, nil
}

func (nameServer *NameServer) DeleteName(_ context.Context, req *pb.DeleteNameRequest) (*pb.DeleteNameResponse, error) {
	id := req.GetId()
	err := nameServer.nameService.DeleteName(id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteNameResponse{
		Message: "Name Deleted Successfully!",
	}, nil
}
