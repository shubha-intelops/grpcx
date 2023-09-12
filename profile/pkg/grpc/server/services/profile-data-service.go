package services

import (
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/daos"
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/models"
)

type ProfileDataService struct {
	profileDataDao *daos.ProfileDataDao
}

func NewProfileDataService() (*ProfileDataService, error) {
	profileDataDao, err := daos.NewProfileDataDao()
	if err != nil {
		return nil, err
	}
	return &ProfileDataService{
		profileDataDao: profileDataDao,
	}, nil
}

func (profileDataService *ProfileDataService) CreateProfileData(profileData *models.ProfileData) (*models.ProfileData, error) {
	return profileDataService.profileDataDao.CreateProfileData(profileData)
}

func (profileDataService *ProfileDataService) UpdateProfileData(id int64, profileData *models.ProfileData) (*models.ProfileData, error) {
	return profileDataService.profileDataDao.UpdateProfileData(id, profileData)
}

func (profileDataService *ProfileDataService) DeleteProfileData(id int64) error {
	return profileDataService.profileDataDao.DeleteProfileData(id)
}

func (profileDataService *ProfileDataService) ListProfileData() ([]*models.ProfileData, error) {
	return profileDataService.profileDataDao.ListProfileData()
}

func (profileDataService *ProfileDataService) GetProfileData(id int64) (*models.ProfileData, error) {
	return profileDataService.profileDataDao.GetProfileData(id)
}
