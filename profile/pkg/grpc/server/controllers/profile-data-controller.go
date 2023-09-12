package controllers

import (
	"context"

	pb "github.com/shubha-intelops/grpcx/profile/gen/api/v1"
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/models"
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/services"
)

type ProfileDataServer struct {
	profileDataService *services.ProfileDataService
	pb.UnimplementedProfileDataServiceServer
}

func NewProfileDataServer() (*ProfileDataServer, error) {
	profileDataService, err := services.NewProfileDataService()
	if err != nil {
		return nil, err
	}
	return &ProfileDataServer{
		profileDataService: profileDataService,
	}, nil
}

func (*ProfileDataServer) Ping(_ context.Context, _ *pb.ProfileDataRequest) (*pb.ProfileDataResponse, error) {
	return &pb.ProfileDataResponse{
		Message: "Server is healthy and working!",
	}, nil
}

func (profileDataServer *ProfileDataServer) CreateProfileData(_ context.Context, req *pb.CreateProfileDataRequest) (*pb.CreateProfileDataResponse, error) {
	profileData := models.ProfileData{

		Address: req.ProfileData.GetAddress(),

		Name: req.ProfileData.GetName(),
	}

	if _, err := profileDataServer.profileDataService.CreateProfileData(&profileData); err != nil {
		return nil, err
	}

	return &pb.CreateProfileDataResponse{
		Message: "ProfileData Created Successfully!",
	}, nil
}

func (profileDataServer *ProfileDataServer) GetProfileData(_ context.Context, req *pb.GetProfileDataRequest) (*pb.GetProfileDataResponse, error) {
	id := req.GetId()
	retrievedProfileData, err := profileDataServer.profileDataService.GetProfileData(id)

	if err != nil {
		return nil, err
	}

	profileDataResponse := &pb.ProfileData{
		Id: id,

		Address: retrievedProfileData.Address,

		Name: retrievedProfileData.Name,
	}

	return &pb.GetProfileDataResponse{
		ProfileData: profileDataResponse,
	}, nil
}

func (profileDataServer *ProfileDataServer) ListProfileData(_ context.Context, _ *pb.ListProfileDataRequest) (*pb.ListProfileDataResponse, error) {
	profileData, err := profileDataServer.profileDataService.ListProfileData()

	if err != nil {
		return nil, err
	}

	var profileDataList []*pb.ProfileData
	for _, retrievedProfileData := range profileData {
		profileDataResponse := &pb.ProfileData{
			Id: retrievedProfileData.Id,

			Address: retrievedProfileData.Address,

			Name: retrievedProfileData.Name,
		}
		profileDataList = append(profileDataList, profileDataResponse)
	}

	return &pb.ListProfileDataResponse{
		ProfileData: profileDataList,
	}, nil
}

func (profileDataServer *ProfileDataServer) UpdateProfileData(_ context.Context, req *pb.UpdateProfileDataRequest) (*pb.UpdateProfileDataResponse, error) {
	id := req.GetId()

	profileDataRequest := models.ProfileData{
		Id: id,

		Address: req.ProfileData.GetAddress(),

		Name: req.ProfileData.GetName(),
	}
	_, err := profileDataServer.profileDataService.UpdateProfileData(id, &profileDataRequest)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateProfileDataResponse{
		Message: "ProfileData Updated Successfully!",
	}, nil
}

func (profileDataServer *ProfileDataServer) DeleteProfileData(_ context.Context, req *pb.DeleteProfileDataRequest) (*pb.DeleteProfileDataResponse, error) {
	id := req.GetId()
	err := profileDataServer.profileDataService.DeleteProfileData(id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteProfileDataResponse{
		Message: "ProfileData Deleted Successfully!",
	}, nil
}
