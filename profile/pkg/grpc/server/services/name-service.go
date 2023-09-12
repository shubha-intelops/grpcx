package services

import (
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/daos"
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/models"
)

type NameService struct {
	nameDao *daos.NameDao
}

func NewNameService() (*NameService, error) {
	nameDao, err := daos.NewNameDao()
	if err != nil {
		return nil, err
	}
	return &NameService{
		nameDao: nameDao,
	}, nil
}

func (nameService *NameService) CreateName(name *models.Name) (*models.Name, error) {
	return nameService.nameDao.CreateName(name)
}

func (nameService *NameService) UpdateName(id int64, name *models.Name) (*models.Name, error) {
	return nameService.nameDao.UpdateName(id, name)
}

func (nameService *NameService) DeleteName(id int64) error {
	return nameService.nameDao.DeleteName(id)
}

func (nameService *NameService) ListNames() ([]*models.Name, error) {
	return nameService.nameDao.ListNames()
}

func (nameService *NameService) GetName(id int64) (*models.Name, error) {
	return nameService.nameDao.GetName(id)
}
