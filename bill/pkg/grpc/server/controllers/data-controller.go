package controllers

import (
	"context"

	pb "github.com/shubha-intelops/grpcx/bill/gen/api/v1"
	"github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/services"
	"github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/models"
)

type DataServer struct {
	dataService *services.DataService
	pb.UnimplementedDataServiceServer
}


func NewDataServer() (*DataServer, error) {
	dataService, err := services.NewDataService()
	if err != nil {
		return nil, err
	}
	return &DataServer{
		dataService: dataService,
	}, nil
}

func (*DataServer) Ping(_ context.Context, _ *pb.DataRequest) (*pb.DataResponse, error) {
	return &pb.DataResponse{
		Message: "Server is healthy and working!",
	}, nil
}

func (dataServer *DataServer) CreateData(_ context.Context, req *pb.CreateDataRequest) (*pb.CreateDataResponse, error) {
	data := models.Data{
        
            
                Ggg: req.Data.GetGgg(),
            
        
	}

	if _, err := dataServer.dataService.CreateData(&data); err != nil {
		return nil, err
	}

	return &pb.CreateDataResponse{
		Message: "Data Created Successfully!",
	}, nil
}


func (dataServer *DataServer) GetData(_ context.Context, req *pb.GetDataRequest) (*pb.GetDataResponse, error) {
    id := req.GetId()
	retrievedData, err := dataServer.dataService.GetData(id)

	if err != nil {
		return nil, err
	}

	dataResponse := &pb.Data{
		Id: id,
        
            
                Ggg:  retrievedData.Ggg,
            
        
	}

	return &pb.GetDataResponse{
		Data: dataResponse,
	}, nil
}

func (dataServer *DataServer) ListData(_ context.Context, _ *pb.ListDataRequest) (*pb.ListDataResponse, error) {
	data, err := dataServer.dataService.ListData()

	if err != nil {
		return nil, err
	}

	var dataList []*pb.Data
	for _, retrievedData := range data {
		dataResponse := &pb.Data{
			Id: retrievedData.Id,
            
                
                    Ggg:  retrievedData.Ggg,
                
            
		}
		dataList = append(dataList, dataResponse)
	}

	return &pb.ListDataResponse{
		Data: dataList,
	}, nil
}

func (dataServer *DataServer) UpdateData(_ context.Context, req *pb.UpdateDataRequest) (*pb.UpdateDataResponse, error) {
	id := req.GetId()

	dataRequest := models.Data{
		Id: id,
        
            
                Ggg:  req.Data.GetGgg(),
            
        
	}
	_, err := dataServer.dataService.UpdateData(id, &dataRequest)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateDataResponse{
		Message: "Data Updated Successfully!",
	}, nil
}

func (dataServer *DataServer) DeleteData(_ context.Context, req *pb.DeleteDataRequest) (*pb.DeleteDataResponse, error) {
	id := req.GetId()
	err := dataServer.dataService.DeleteData(id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteDataResponse{
		Message: "Data Deleted Successfully!",
	}, nil
}