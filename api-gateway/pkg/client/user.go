package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"

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
		Otp:   uint(res.Otp),
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
			Age:       uint(u.Age),
			GenderID:  uint(u.GenderId),
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
func (c *userClient) AddUserInterest(userID uint64, interestID int) error {
	_, err := c.Client.AddUserInterest(context.Background(), &pb.AddUserInterestRequest{
		UserId:     userID,
		InterestId: uint64(interestID),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *userClient) EditUserInterest(userID uint64, interestID uint64, newInterestName string) error {
	_, err := c.Client.EditUserInterest(context.Background(), &pb.EditUserInterestRequest{
		UserId:          userID,
		InterestId:      interestID,
		NewInterestName: newInterestName,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *userClient) DeleteUserInterest(userID uint64, interestID uint64) error {
	_, err := c.Client.DeleteUserInterest(context.Background(), &pb.DeleteUserInterestRequest{
		UserId:     userID,
		InterestId: interestID,
	})
	if err != nil {
		return err
	}
	return nil
}
func (c *userClient) AddUserPreference(userID uint64, preference_id int) error {
	_, err := c.Client.AddUserPreference(context.Background(), &pb.AddUserPreferenceRequest{
		UserId:       userID,
		PreferenceId: uint64(preference_id),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *userClient) EditUserPreference(userID uint64, preferenceID uint64, newPreferenceName string) error {
	_, err := c.Client.EditUserPreference(context.Background(), &pb.EditUserPreferenceRequest{
		UserId:            userID,
		PreferenceId:      preferenceID,
		NewPreferenceName: newPreferenceName,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *userClient) DeleteUserPreference(userID uint64, preferenceID uint64) error {
	_, err := c.Client.DeleteUserPreference(context.Background(), &pb.DeleteUserPreferenceRequest{
		UserId:       userID,
		PreferenceId: preferenceID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *userClient) GetUserInterests(userID uint64) ([]string, error) {
	interestsResponse, err := c.Client.GetUserInterests(context.Background(), &pb.GetUserInterestsRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	return interestsResponse.Interests, nil
}

func (c *userClient) GetUserPreferences(userID uint64) ([]string, error) {
	preferencesResponse, err := c.Client.GetUserPreferences(context.Background(), &pb.GetUserPreferencesRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	return preferencesResponse.Preferences, nil
}

func (c *userClient) FollowUser(userID int, senderID int) error {
	_, err := c.Client.FollowUser(context.Background(), &pb.FollowUserRequest{
		Userid:   int64(userID),
		Senderid: int64(senderID),
	})
	if err != nil {
		return err
	}

	return nil
}
func (c *userClient) BlockUser(userID int, senderID int) error {
	_, err := c.Client.BlockUser(context.Background(), &pb.BlockUserRequest{
		Userid:   int64(userID),
		Senderid: int64(senderID),
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *userClient) SendMessage(message models.UserMessage) (models.UserMessage, error) {
	var mediaList []*pb.Media
	for _, m := range message.Media {
		mediaList = append(mediaList, &pb.Media{
			Filename: m.Filename,
		})
	}
	req := &pb.SendMessageRequest{
		SenterId:   uint32(message.SenderID),
		ReceiverId: uint32(message.RecipentID),
		Content:    message.Content,
		Media:      mediaList, // Assign []*room.Media
	}

	res, err := c.Client.SendMessage(context.Background(), req)
	if err != nil {
		return models.UserMessage{}, err
	}
	return models.UserMessage{
		ID:         uint(res.MessageId),
		SenderID:   uint(res.SenterId),
		RecipentID: uint(res.ReceiverId),
		Content:    res.Content,
		CreatedAt:  res.Timestamp.AsTime(), // Convert Protobuf timestamp to Go time
	}, nil
}
func (c *userClient) ReadMessages(userid uint32, senterid uint32) ([]models.UserMessage, error) {
	req := &pb.ReadMessagesRequest{
		ReceiverId: userid,
		SenterId:   senterid,
	}
	res, err := c.Client.ReadMessages(context.Background(), req)
	if err != nil {
		return nil, err
	}
	messages := make([]models.UserMessage, len(res.Messages))
	for i, pbMsg := range res.Messages {
		messages[i] = models.UserMessage{
			ID:        uint(pbMsg.MessageId),
			SenderID:  uint(pbMsg.SenderId),
			Content:   pbMsg.Content,
			CreatedAt: pbMsg.Timestamp.AsTime(), // Convert Protobuf timestamp to Go time
		}
	}
	return messages, nil
}
func (c *userClient) GetConnections(userID uint64) ([]models.UserDetail, error) {
	// Send gRPC request to server
	response, err := c.Client.GetConnections(context.Background(), &pb.GetConnectionsRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	// Convert []*pb.User to []models.User
	var users []models.UserDetail
	for _, u := range response.UserDetails {
		user := models.UserDetail{
			ID:        uint(u.Id),
			Firstname: u.Firstname,
			Lastname:  u.Lastname,
			Email:     u.Email,
		}
		users = append(users, user)
	}

	return users, nil
}
func (c *userClient) UpdateProfilePhoto(ID int64, files []*multipart.FileHeader) error {
	var imagePaths [][]byte

	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		buffer := new(bytes.Buffer)
		if _, err := io.Copy(buffer, f); err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}

		imagePaths = append(imagePaths, buffer.Bytes())
	}

	// Convert user model to gRPC request
	grpcReq := &pb.UpdateProfilePhotoRequest{
		Userid:    ID,
		ImageData: imagePaths,
	}

	// Call gRPC client function
	_, err := c.Client.UpdateProfilePhoto(context.Background(), grpcReq)
	if err != nil {
		return fmt.Errorf("failed to update profile photo: %w", err)
	}
	return nil
}

// func (c *userClient) AddProfilePhoto(ID int64, file *multipart.FileHeader) error {
// 	// Open the file
// 	f, err := file.Open()
// 	if err != nil {
// 		return fmt.Errorf("failed to open file: %w", err)
// 	}
// 	defer f.Close()

// 	// Read the file into a buffer
// 	buffer := new(bytes.Buffer)
// 	if _, err := io.Copy(buffer, f); err != nil {
// 		return fmt.Errorf("failed to read file: %w", err)
// 	}

// 	// Convert user model to gRPC request
// 	grpcReq := &pb.UpdateProfilePhotoRequest{
// 		Userid:    ID,
// 		ImageData: buffer.Bytes(),
// 	}

// 	// Call gRPC client function
// 	_, err = c.Client.UpdateProfilePhoto(context.Background(), grpcReq)
// 	if err != nil {
// 		return fmt.Errorf("failed to update profile photo: %w", err)
// 	}
// 	return nil
// }
