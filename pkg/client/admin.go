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
func (ad *adminClient) GetPreferences() ([]models.Intrests, error) {
	usersResponse, err := ad.Client.GetPreferences(context.Background(), &pb.GetPreferencesRequest{})
	if err != nil {
		return nil, err
	}

	// Convert []*pb.Users to []models.Users
	var Preferences []models.Intrests
	for _, u := range usersResponse.Preferences {
		PreferencesModel := models.Intrests{
			ID:   int(u.Id),
			Name: u.PreferenceName,
		}
		Preferences = append(Preferences, PreferencesModel)
	}

	return Preferences, nil
}
