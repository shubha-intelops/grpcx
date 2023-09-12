package services

import (
	"github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/daos"
	"github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/models"
)

type BillService struct {
	billDao *daos.BillDao
}

func NewBillService() (*BillService, error) {
	billDao, err := daos.NewBillDao()
	if err != nil {
		return nil, err
	}
	return &BillService{
		billDao: billDao,
	}, nil
}

func (billService *BillService) CreateBill(bill *models.Bill) (*models.Bill, error) {
	return billService.billDao.CreateBill(bill)
}

func (billService *BillService) UpdateBill(id int64, bill *models.Bill) (*models.Bill, error) {
	return billService.billDao.UpdateBill(id, bill)
}

func (billService *BillService) DeleteBill(id int64) error {
	return billService.billDao.DeleteBill(id)
}

func (billService *BillService) ListBills() ([]*models.Bill, error) {
	return billService.billDao.ListBills()
}

func (billService *BillService) GetBill(id int64) (*models.Bill, error) {
	return billService.billDao.GetBill(id)
}
