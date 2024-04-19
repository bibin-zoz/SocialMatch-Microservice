package client

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/config"
	pb "github.com/bibin-zoz/api-gateway/pkg/pb/userauth"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userClient struct {
	Client pb.UserClient
}

func NewUserClient(cfg config.Config) interfaces.UserClient {
	fmt.Println("sasa")

	grpcConnection, err := grpc.NewClient(cfg.UserSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect", err)
		return nil
	}

	grpcClient := pb.NewUserClient(grpcConnection)
	fmt.Println("grpc", grpcClient)
	return &userClient{
		Client: grpcClient,
	}

}
func (c *userClient) UsersSignUp(user models.UserSignup) (models.TokenUser, error) {

	res, err := c.Client.UserSignUp(context.Background(), &pb.UserSignUpRequest{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     fmt.Sprint(user.Number),
	})
	if err != nil {
		return models.TokenUser{}, err
	}
	userDetails := models.UserDetails{
		ID:        uint(res.UserDetails.Id),
		Firstname: res.UserDetails.Firstname,
		Lastname:  res.UserDetails.Lastname,
		Email:     res.UserDetails.Email,
		Phone:     res.UserDetails.Phone,
	}

	return models.TokenUser{
		User:         userDetails,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
func (c *userClient) UserLogin(user models.UserLogin) (models.TokenUser, error) {
	res, err := c.Client.UserLogin(context.Background(), &pb.UserLoginRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		fmt.Println("🤷‍♂️")
		return models.TokenUser{}, err
	}
	userDetails := models.UserDetails{
		ID:        uint(res.UserDetails.Id),
		Firstname: res.UserDetails.Firstname,
		Lastname:  res.UserDetails.Lastname,
		Email:     res.UserDetails.Email,
		Phone:     res.UserDetails.Phone,
	}

	return models.TokenUser{
		User:         userDetails,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
func (c *userClient) UserEditDetails(user models.UserUpdateDetails) (models.UserDetails, error) {

	res, err := c.Client.UserEditDetails(context.Background(), &pb.UserEditDetailsRequest{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Username:  user.Username,
	})
	if err != nil {
		return models.UserDetails{}, err
	}
	userDetails := models.UserDetails{
		ID:        uint(res.UserDetails.Id),
		Firstname: res.UserDetails.Firstname,
		Lastname:  res.UserDetails.Lastname,
		Email:     res.UserDetails.Email,
		Phone:     res.UserDetails.Phone,
	}

	return userDetails, nil
}
func (c *userClient) UserOtpRequest(user models.UserVerificationRequest) (models.Otp, error) {
	res, err := c.Client.UserOtpGeneration(context.Background(), &pb.UserOtpRequest{
		Email: user.Email,
	})
	if err != nil {
		return models.Otp{}, err
	}
	return models.Otp{
		Email: res.Email,
		Otp:   int(res.Otp),
	}, nil

}

func (c *userClient) UserOtpVerificationReq(user models.Otp) (models.UserDetail, error) {
	_, err := c.Client.UserOtpVerification(context.Background(), &pb.UserOtpVerificationRequest{
		Email: user.Email,
		Otp:   int64(user.Otp),
	})
	if err != nil {
		return models.UserDetail{}, err
	}
	return models.UserDetail{}, nil

}

func (c *userClient) GetAllUsers() ([]models.Users, error) {
	usersResponse, err := c.Client.GetUsers(context.Background(), &pb.GetUsersRequest{})
	if err != nil {
		return nil, err
	}

	// Convert []*pb.Users to []models.Users
	var users []models.Users
	for _, u := range usersResponse.Users {
		userModel := models.Users{
			ID:        uint(u.Id),
			Username:  u.Username,
			Email:     u.Email,
			Firstname: u.Firstname,
			Lastname:  u.Lastname,
			Phone:     u.Phone,
			Password:  "hidden",
			Age:       int(u.Age),
			GenderID:  int(u.GenderId),
			Blocked:   u.Blocked,
		}
		users = append(users, userModel)
	}

	return users, nil
}

func (c *userClient) UpdateUserStatus(userid int64) (models.UserDetail, error) {
	fmt.Println("updateuserstaus cleint code")
	_, err := c.Client.UpdateUserStatus(context.Background(), &pb.UpdateUserStatusRequest{
		Userid: userid,
	})
	if err != nil {
		return models.UserDetail{}, err
	}
	return models.UserDetail{}, nil
}
