package services

import (
	"github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/daos"
	"github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/models"
)

type DataService struct {
	dataDao *daos.DataDao
}

func NewDataService() (*DataService, error) {
	dataDao, err := daos.NewDataDao()
	if err != nil {
		return nil, err
	}
	return &DataService{
		dataDao: dataDao,
	}, nil
}

func (dataService *DataService) CreateData(data *models.Data) (*models.Data, error)  {
	return dataService.dataDao.CreateData(data)
}

func (dataService *DataService) UpdateData(id int64, data *models.Data) (*models.Data, error)  {
	return dataService.dataDao.UpdateData(id, data)
}

func (dataService *DataService) DeleteData(id int64) error {
	return dataService.dataDao.DeleteData(id)
}

func (dataService *DataService) ListData() ([]*models.Data, error) {
	return dataService.dataDao.ListData()
}

func (dataService *DataService) GetData(id int64) (*models.Data, error) {
	return dataService.dataDao.GetData(id)
}
