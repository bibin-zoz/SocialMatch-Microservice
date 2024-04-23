package client

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/config"
	pb "github.com/bibin-zoz/api-gateway/pkg/pb/admin"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type adminClient struct {
	Client pb.AdminClient
}

func NewAdminClient(cfg config.Config) interfaces.AdminClient {

	grpcConnection, err := grpc.Dial(cfg.AdminSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewAdminClient(grpcConnection)

	return &adminClient{
		Client: grpcClient,
	}

}

func (ad *adminClient) AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminLogin(context.Background(), &pb.AdminLoginRequest{
		Email:    adminDetails.Email,
		Password: adminDetails.Password,
	})

	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Token: admin.Token,
	}, nil
}
func (ad *adminClient) GetIntrests() ([]models.Intrests, error) {
	usersResponse, err := ad.Client.GetInterests(context.Background(), &pb.GetInterestsRequest{})
	if err != nil {
		return nil, err
	}

	// Convert []*pb.Users to []models.Users
	var Intrests []models.Intrests
	for _, u := range usersResponse.Interests {
		userModel := models.Intrests{
			ID:   int(u.Id),
			Name: u.InterestName,
		}
		Intrests = append(Intrests, userModel)
	}

	return Intrests, nil
}
func (ad *adminClient) GetPreferences() ([]models.Preferences, error) {
	usersResponse, err := ad.Client.GetPreferences(context.Background(), &pb.GetPreferencesRequest{})
	if err != nil {
		return nil, err
	}

	// Convert []*pb.Users to []models.Users
	var Preferences []models.Preferences
	for _, u := range usersResponse.Preferences {
		PreferencesModel := models.Preferences{
			ID:   int(u.Id),
			Name: u.PreferenceName,
		}
		Preferences = append(Preferences, PreferencesModel)
	}

	return Preferences, nil
}
func (ad *adminClient) AddInterest(interestName string) (models.InterestResponse, error) {
	response, err := ad.Client.AddInterest(context.Background(), &pb.AddInterestRequest{
		InterestName: interestName,
	})
	fmt.Println("interestName client", interestName)
	if err != nil {
		return models.InterestResponse{}, err
	}
	return models.InterestResponse{
		ID: int(response.Id),
	}, nil
}

func (ad *adminClient) EditInterest(interestID int, newInterestName string) (models.InterestResponse, error) {
	_, err := ad.Client.EditInterest(context.Background(), &pb.EditInterestRequest{
		Id:           int64(interestID),
		InterestName: newInterestName,
	})
	if err != nil {
		return models.InterestResponse{}, err
	}
	return models.InterestResponse{
		ID:   int(interestID),
		Name: newInterestName,
	}, nil
}

func (ad *adminClient) DeleteInterest(interestID int) error {
	_, err := ad.Client.DeleteInterest(context.Background(), &pb.DeleteInterestRequest{
		Id: int64(interestID),
	})
	return err
}

func (ad *adminClient) AddPreference(preferenceName string) (models.PreferenceResponse, error) {
	response, err := ad.Client.AddPreference(context.Background(), &pb.AddPreferenceRequest{
		PreferenceName: preferenceName,
	})
	if err != nil {
		return models.PreferenceResponse{}, err
	}
	return models.PreferenceResponse{
		ID:   int(response.Id),
		Name: preferenceName,
	}, nil
}

func (ad *adminClient) EditPreference(preferenceID int, newPreferenceName string) (models.PreferenceResponse, error) {
	_, err := ad.Client.EditPreference(context.Background(), &pb.EditPreferenceRequest{
		Id:             int64(preferenceID),
		PreferenceName: newPreferenceName,
	})
	if err != nil {
		return models.PreferenceResponse{}, err
	}
	return models.PreferenceResponse{
		ID:   int(preferenceID),
		Name: newPreferenceName,
	}, nil
}

func (ad *adminClient) DeletePreference(preferenceID int) error {
	_, err := ad.Client.DeletePreference(context.Background(), &pb.DeletePreferenceRequest{
		Id: int64(preferenceID),
	})
	return err
}
