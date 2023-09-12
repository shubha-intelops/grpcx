package controllers

import (
	"context"

	pb "github.com/shubha-intelops/grpcx/bill/gen/api/v1"
	"github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/services"
	"github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/models"
)

type BillServer struct {
	billService *services.BillService
	pb.UnimplementedBillServiceServer
}


func NewBillServer() (*BillServer, error) {
	billService, err := services.NewBillService()
	if err != nil {
		return nil, err
	}
	return &BillServer{
		billService: billService,
	}, nil
}

func (*BillServer) Ping(_ context.Context, _ *pb.BillRequest) (*pb.BillResponse, error) {
	return &pb.BillResponse{
		Message: "Server is healthy and working!",
	}, nil
}

func (billServer *BillServer) CreateBill(_ context.Context, req *pb.CreateBillRequest) (*pb.CreateBillResponse, error) {
	bill := models.Bill{
        
            
                BillAdd: req.Bill.GetBillAdd(),
            
        
            
                Rate: req.Bill.GetRate(),
            
        
	}

	if _, err := billServer.billService.CreateBill(&bill); err != nil {
		return nil, err
	}

	return &pb.CreateBillResponse{
		Message: "Bill Created Successfully!",
	}, nil
}


func (billServer *BillServer) GetBill(_ context.Context, req *pb.GetBillRequest) (*pb.GetBillResponse, error) {
    id := req.GetId()
	retrievedBill, err := billServer.billService.GetBill(id)

	if err != nil {
		return nil, err
	}

	billResponse := &pb.Bill{
		Id: id,
        
            
                BillAdd:  retrievedBill.BillAdd,
            
        
            
                Rate:  retrievedBill.Rate,
            
        
	}

	return &pb.GetBillResponse{
		Bill: billResponse,
	}, nil
}

func (billServer *BillServer) ListBills(_ context.Context, _ *pb.ListBillsRequest) (*pb.ListBillsResponse, error) {
	bills, err := billServer.billService.ListBills()

	if err != nil {
		return nil, err
	}

	var billList []*pb.Bill
	for _, retrievedBill := range bills {
		billResponse := &pb.Bill{
			Id: retrievedBill.Id,
            
                
                    BillAdd:  retrievedBill.BillAdd,
                
            
                
                    Rate:  retrievedBill.Rate,
                
            
		}
		billList = append(billList, billResponse)
	}

	return &pb.ListBillsResponse{
		Bill: billList,
	}, nil
}

func (billServer *BillServer) UpdateBill(_ context.Context, req *pb.UpdateBillRequest) (*pb.UpdateBillResponse, error) {
	id := req.GetId()

	billRequest := models.Bill{
		Id: id,
        
            
                BillAdd:  req.Bill.GetBillAdd(),
            
        
            
                Rate:  req.Bill.GetRate(),
            
        
	}
	_, err := billServer.billService.UpdateBill(id, &billRequest)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateBillResponse{
		Message: "Bill Updated Successfully!",
	}, nil
}

func (billServer *BillServer) DeleteBill(_ context.Context, req *pb.DeleteBillRequest) (*pb.DeleteBillResponse, error) {
	id := req.GetId()
	err := billServer.billService.DeleteBill(id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteBillResponse{
		Message: "Bill Deleted Successfully!",
	}, nil
}