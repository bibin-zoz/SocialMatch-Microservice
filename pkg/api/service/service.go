package service

import (
	"context"
	"fmt"

	pb "github.com/bibin-zoz/social-match-admin-svc/pkg/pb/admin"
	interfaces "github.com/bibin-zoz/social-match-admin-svc/pkg/usecase/interface"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/utils/models"
)

type AdminServer struct {
	adminUseCase interfaces.AdminUseCase
	pb.UnimplementedAdminServer
}

func NewAdminServer(useCase interfaces.AdminUseCase) pb.AdminServer {

	return &AdminServer{
		adminUseCase: useCase,
	}

}

func (ad *AdminServer) AdminLogin(ctx context.Context, Req *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	adminLogin := models.AdminLogin{
		Email:    Req.Email,
		Password: Req.Password,
	}
	fmt.Println("admin login", adminLogin)
	admin, err := ad.adminUseCase.LoginHandler(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{}, err
	}
	fmt.Println("true")

	return &pb.AdminLoginResponse{
		Status: 200,
		Token:  admin.Token,
	}, nil
}
func (ad *AdminServer) GetInterests(ctx context.Context, req *pb.GetInterestsRequest) (*pb.GetInterestsResponse, error) {
	Interest, err := ad.adminUseCase.GetInterests()
	if err != nil {
		return nil, err
	}

	var pbInterest []*pb.Interest
	for _, u := range Interest {
		pbRes := &pb.Interest{
			Id:           (int64(u.ID)),
			InterestName: u.InterestName,
		}
		pbInterest = append(pbInterest, pbRes)
	}

	return &pb.GetInterestsResponse{
		Status:    201,
		Interests: pbInterest,
	}, nil
}
func (ad *AdminServer) AddInterest(ctx context.Context, req *pb.AddInterestRequest) (*pb.AddInterestResponse, error) {
	interestName := req.InterestName
	// Call the corresponding use case method to add interest
	id, err := ad.adminUseCase.AddInterest(interestName)
	if err != nil {
		return nil, err
	}
	return &pb.AddInterestResponse{
		Status: 200,
		Id:     id,
	}, nil
}

func (ad *AdminServer) EditInterest(ctx context.Context, req *pb.EditInterestRequest) (*pb.EditInterestResponse, error) {
	id := req.Id
	interestName := req.InterestName
	// Call the corresponding use case method to edit interest
	err := ad.adminUseCase.EditInterest(id, interestName)
	if err != nil {
		return nil, err
	}
	return &pb.EditInterestResponse{
		Status: 200,
	}, nil
}

func (ad *AdminServer) DeleteInterest(ctx context.Context, req *pb.DeleteInterestRequest) (*pb.DeleteInterestResponse, error) {
	id := req.Id
	// Call the corresponding use case method to delete interest
	err := ad.adminUseCase.DeleteInterest(id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteInterestResponse{
		Status: 200,
	}, nil
}

func (ad *AdminServer) AddPreference(ctx context.Context, req *pb.AddPreferenceRequest) (*pb.AddPreferenceResponse, error) {
	preferenceName := req.PreferenceName
	// Call the corresponding use case method to add preference
	id, err := ad.adminUseCase.AddPreference(preferenceName)
	if err != nil {
		return nil, err
	}
	return &pb.AddPreferenceResponse{
		Status: 200,
		Id:     id,
	}, nil
}

func (ad *AdminServer) EditPreference(ctx context.Context, req *pb.EditPreferenceRequest) (*pb.EditPreferenceResponse, error) {
	id := req.Id
	preferenceName := req.PreferenceName
	// Call the corresponding use case method to edit preference
	err := ad.adminUseCase.EditPreference(id, preferenceName)
	if err != nil {
		return nil, err
	}
	return &pb.EditPreferenceResponse{
		Status: 200,
	}, nil
}

func (ad *AdminServer) DeletePreference(ctx context.Context, req *pb.DeletePreferenceRequest) (*pb.DeletePreferenceResponse, error) {
	id := req.Id
	// Call the corresponding use case method to delete preference
	err := ad.adminUseCase.DeletePreference(id)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePreferenceResponse{
		Status: 200,
	}, nil
}
func (ad *AdminServer) GetPreferences(ctx context.Context, req *pb.GetPreferencesRequest) (*pb.GetPreferencesResponse, error) {
	preference, err := ad.adminUseCase.GetPreferences()
	if err != nil {
		return nil, err
	}

	var pbPreferences []*pb.Preference
	for _, u := range preference {
		pbRes := &pb.Preference{
			Id:             (int64(u.ID)),
			PreferenceName: u.PreferenceName,
		}
		pbPreferences = append(pbPreferences, pbRes)
	}

	return &pb.GetPreferencesResponse{
		Status:      201,
		Preferences: pbPreferences,
	}, nil
}
