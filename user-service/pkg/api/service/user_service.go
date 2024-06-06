package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"

	interfaces "github.com/bibin-zoz/social-match-userauth-svc/pkg/usecase/interface"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
)

type UserServer struct {
	userUseCase interfaces.UserUseCase
	pb.UnimplementedUserServer
}

func NewUserServer(useCase interfaces.UserUseCase) pb.UserServer {

	return &UserServer{
		userUseCase: useCase,
	}

}
func (a *UserServer) UserSignUp(ctx context.Context, userSignUpDetails *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {

	userCreateDetails := models.UserSignup{
		FirstName: userSignUpDetails.Firstname,
		LastName:  userSignUpDetails.Lastname,
		Email:     userSignUpDetails.Email,
		Number:    userSignUpDetails.Phone,
		Password:  userSignUpDetails.Password,
		Age:       int(userSignUpDetails.Age),
		GenderID:  int(userSignUpDetails.Genderid),
	}

	data, err := a.userUseCase.UsersSignUp(userCreateDetails)
	if err != nil {
		return &pb.UserSignUpResponse{}, err
	}
	userDetails := &pb.UserDetails{Id: uint64(data.User.ID), Firstname: data.User.Firstname, Lastname: data.User.Lastname, Email: data.User.Email, Phone: data.User.Phone}
	return &pb.UserSignUpResponse{
		Status:       201,
		UserDetails:  userDetails,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}, nil

}
func (a *UserServer) UserLogin(ctx context.Context, loginDeatails *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	login := models.UserLogin{
		Email:    loginDeatails.Email,
		Password: loginDeatails.Password,
	}
	data, err := a.userUseCase.UsersLogin(login)
	if err != nil {
		return &pb.UserLoginResponse{}, err
	}
	userDetails := &pb.UserDetails{
		Id:        uint64(data.User.ID),
		Firstname: data.User.Firstname,
		Lastname:  data.User.Lastname,
		Email:     data.User.Email,
		Phone:     data.User.Phone,
	}
	return &pb.UserLoginResponse{
		Status:       201,
		UserDetails:  userDetails,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}, nil
}

func (a *UserServer) UserEditDetails(ctx context.Context, userEditDetails *pb.UserEditDetailsRequest) (*pb.UserEditDetailsResponse, error) {

	userCreateDetails := models.UserSignup{
		FirstName: userEditDetails.Firstname,
		LastName:  userEditDetails.Lastname,
		Email:     userEditDetails.Email,
		Number:    userEditDetails.Phone,
		Password:  userEditDetails.Password,
	}

	user, err := a.userUseCase.UserEditDetails(userCreateDetails)
	if err != nil {
		return &pb.UserEditDetailsResponse{}, err
	}
	userDetails := &pb.UserDetails{Id: uint64(user.ID), Firstname: user.Firstname, Lastname: user.Lastname, Email: user.Email, Phone: user.Phone}
	return &pb.UserEditDetailsResponse{
		Status:      201,
		UserDetails: userDetails,
	}, nil

}

func (a *UserServer) UserOtpGeneration(ctx context.Context, userEditDetails *pb.UserOtpRequest) (*pb.UserOtpRequestResponse, error) {

	email := userEditDetails.Email

	res, err := a.userUseCase.UserGenerateOtp(email)
	otp, _ := strconv.Atoi(res)
	if err != nil {
		return &pb.UserOtpRequestResponse{}, err
	}
	return &pb.UserOtpRequestResponse{
		Status: 201,
		Email:  email,
		Otp:    int64(otp),
	}, nil

}
func (a *UserServer) UserOtpVerification(ctx context.Context, req *pb.UserOtpVerificationRequest) (*pb.UserOtpVerificationResponse, error) {

	_, err := a.userUseCase.UserVerifyOtp(fmt.Sprint(req.Otp), req.Email)
	if err != nil {
		return &pb.UserOtpVerificationResponse{}, err
	}

	return &pb.UserOtpVerificationResponse{
		Status:      201,
		UserDetails: nil,
	}, nil

}

func (a *UserServer) AddUserInterest(ctx context.Context, req *pb.AddUserInterestRequest) (*pb.AddUserInterestResponse, error) {
	err := a.userUseCase.AddUserInterest(req.UserId, int(req.InterestId))
	if err != nil {
		return &pb.AddUserInterestResponse{}, err
	}
	return &pb.AddUserInterestResponse{Status: 201}, nil
}

func (a *UserServer) EditUserInterest(ctx context.Context, req *pb.EditUserInterestRequest) (*pb.EditUserInterestResponse, error) {
	err := a.userUseCase.EditUserInterest(req.UserId, req.InterestId, req.NewInterestName)
	if err != nil {
		return &pb.EditUserInterestResponse{}, err
	}
	return &pb.EditUserInterestResponse{Status: 201}, nil
}

func (a *UserServer) DeleteUserInterest(ctx context.Context, req *pb.DeleteUserInterestRequest) (*pb.DeleteUserInterestResponse, error) {
	err := a.userUseCase.DeleteUserInterest(req.UserId, req.InterestId)
	if err != nil {
		return &pb.DeleteUserInterestResponse{}, err
	}
	return &pb.DeleteUserInterestResponse{Status: 201}, nil
}
func (a *UserServer) AddUserPreference(ctx context.Context, req *pb.AddUserPreferenceRequest) (*pb.AddUserPreferenceResponse, error) {
	err := a.userUseCase.AddUserPreference(req.UserId, int(req.PreferenceId))
	if err != nil {
		return &pb.AddUserPreferenceResponse{}, err
	}
	return &pb.AddUserPreferenceResponse{Status: 201}, nil
}

func (a *UserServer) EditUserPreference(ctx context.Context, req *pb.EditUserPreferenceRequest) (*pb.EditUserPreferenceResponse, error) {
	err := a.userUseCase.EditUserPreference(req.UserId, req.PreferenceId, req.NewPreferenceName)
	if err != nil {
		return &pb.EditUserPreferenceResponse{}, err
	}
	return &pb.EditUserPreferenceResponse{Status: 201}, nil
}

func (a *UserServer) DeleteUserPreference(ctx context.Context, req *pb.DeleteUserPreferenceRequest) (*pb.DeleteUserPreferenceResponse, error) {
	err := a.userUseCase.DeleteUserPreference(req.UserId, req.PreferenceId)
	if err != nil {
		return &pb.DeleteUserPreferenceResponse{}, err
	}
	return &pb.DeleteUserPreferenceResponse{Status: 201}, nil
}

func (a *UserServer) GetUserInterests(ctx context.Context, req *pb.GetUserInterestsRequest) (*pb.GetUserInterestsResponse, error) {
	interests, err := a.userUseCase.GetUserInterests(req.UserId)
	if err != nil {
		return &pb.GetUserInterestsResponse{}, err
	}
	return &pb.GetUserInterestsResponse{
		Status:    201,
		Interests: interests,
	}, nil
}

func (a *UserServer) GetUserPreferences(ctx context.Context, req *pb.GetUserPreferencesRequest) (*pb.GetUserPreferencesResponse, error) {
	preferences, err := a.userUseCase.GetUserPreferences(req.UserId)
	if err != nil {
		return &pb.GetUserPreferencesResponse{}, err
	}
	return &pb.GetUserPreferencesResponse{
		Status:      201,
		Preferences: preferences,
	}, nil
}
func (a *UserServer) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.FollowUserResponce, error) {
	fmt.Println("follow user service")
	err := a.userUseCase.FollowUser(req.Senderid, req.Userid)
	if err != nil {
		return &pb.FollowUserResponce{}, err
	}
	return &pb.FollowUserResponce{
		Status: 201,
	}, nil
}
func (a *UserServer) BlockUser(ctx context.Context, req *pb.BlockUserRequest) (*pb.BlockUserResponce, error) {
	err := a.userUseCase.BlockConnection(req.Senderid, req.Userid)
	if err != nil {
		return &pb.BlockUserResponce{}, err
	}
	return &pb.BlockUserResponce{
		Status: 201,
	}, nil
}
func (a *UserServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponce, error) {
	media := make([]models.Media, len(req.Media))
	for i, m := range req.Media {
		media[i] = models.Media{
			Filename: m.Filename,
		}
	}

	message := &models.UserMessage{
		SenderID:   uint(req.SenterId),
		RecipentID: uint(req.ReceiverId),
		Content:    req.Content,
		CreatedAt:  time.Now(),
		Media:      media,
	}
	err := a.userUseCase.SendMessage(message)
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponce{
		MessageId:  123,
		SenterId:   req.SenterId,
		ReceiverId: req.ReceiverId,
		Content:    req.Content,
		Timestamp:  &timestamp.Timestamp{},
		Media:      req.Media,
	}, nil
}
func (a *UserServer) ReadMessages(ctx context.Context, req *pb.ReadMessagesRequest) (*pb.ReadMessagesResponse, error) {
	// Call the use case method to get messages for the specified room
	messages, err := a.userUseCase.GetMessages(uint64(req.ReceiverId), uint64(req.SenterId))
	if err != nil {
		return nil, err
	}

	// Convert the returned messages to protobuf format and return
	var pbMessages []*pb.Message
	for _, message := range messages {
		timestamp := timestamppb.New(message.CreatedAt)
		pbMessage := &pb.Message{
			MessageId:  uint32(message.ID),
			ReceiverId: uint32(message.RecipentID),
			SenderId:   uint32(message.SenderID),
			Content:    message.Content,
			Timestamp:  timestamp,
		}
		pbMessages = append(pbMessages, pbMessage)
	}

	return &pb.ReadMessagesResponse{
		Messages: pbMessages,
	}, nil
}
func (a *UserServer) GetConnections(ctx context.Context, req *pb.GetConnectionsRequest) (*pb.GetConnectionsResponse, error) {
	// Call the use case method to get connections for the specified user ID
	connections, err := a.userUseCase.GetConnections(req.UserId)
	if err != nil {
		return nil, err
	}

	// Convert the returned connections to protobuf format and return
	var pbConnections []*pb.UserDetails
	for _, connection := range connections {
		pbConnection := &pb.UserDetails{
			Id:        uint64(connection.ID),
			Firstname: connection.Firstname,
			Lastname:  connection.Lastname,
			Email:     connection.Email,
			Phone:     connection.Phone,
			// Populate other fields as needed
		}
		pbConnections = append(pbConnections, pbConnection)
	}
	fmt.Println("userdetails", pbConnections)
	return &pb.GetConnectionsResponse{
		Status:      201,
		UserDetails: pbConnections,
	}, nil
}
func (s *UserServer) UpdateProfilePhoto(ctx context.Context, req *pb.UpdateProfilePhotoRequest) (*pb.UpdateUserResponse, error) {
	userPics := &models.UserProfilePhoto{
		UserID:    int(req.Userid),
		ImageData: req.ImageData,
	}
	res, err := s.userUseCase.UpdateProfilePhoto(*userPics)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		Status:   0,
		Userid:   int64(res.UserID),
		ImageUrl: res.ImageURL,
	}, nil
}
func (s *UserServer) AddProfilePhoto(ctx context.Context, req *pb.AddProfilePhotoRequest) (*pb.AddProfilePhotoResponse, error) {
	// userPics := &models.UserProfilePhoto{
	// 	UserID:    int(req.Userid),
	// 	ImageData: req.ImageData,
	// }
	_, err := s.userUseCase.AddProfilePhoto(uint32(req.Userid), req.ImageData)
	if err != nil {
		return nil, err
	}
	return &pb.AddProfilePhotoResponse{
		Status:   201,
		ImageUrl: "",
	}, nil
}

func (s *UserServer) DeleteProfilePhotoByID(ctx context.Context, req *pb.DeleteProfilePhotoRequest) (*pb.DeleteProfilePhotoResponse, error) {
	_, err := s.userUseCase.DeleteProfilePhotoByID(uint32(req.Userid), req.ImageId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProfilePhotoResponse{
		Status: 0,
	}, nil
}
