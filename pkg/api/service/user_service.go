package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/pb"

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
	// Convert request parameters to domain model
	interest := models.Interest{
		// Add necessary fields from req to interest model
	}

	// Call use case method to add user interest
	err := a.userUseCase.AddUserInterest(req.UserId, interest)
	if err != nil {
		return &pb.AddUserInterestResponse{}, err
	}

	return &pb.AddUserInterestResponse{
		Status: 201,
	}, nil
}

func (a *UserServer) EditUserInterest(ctx context.Context, req *pb.EditUserInterestRequest) (*pb.EditUserInterestResponse, error) {
	// Convert request parameters to domain model
	interest := models.Interest{
		// Add necessary fields from req to interest model
	}

	// Call use case method to edit user interest
	err := a.userUseCase.EditUserInterest(req.UserId, req.InterestId, interest)
	if err != nil {
		return &pb.EditUserInterestResponse{}, err
	}

	return &pb.EditUserInterestResponse{
		Status: 201,
	}, nil
}

func (a *UserServer) DeleteUserInterest(ctx context.Context, req *pb.DeleteUserInterestRequest) (*pb.DeleteUserInterestResponse, error) {
	// Call use case method to delete user interest
	err := a.userUseCase.DeleteUserInterest(req.UserId, req.InterestId)
	if err != nil {
		return &pb.DeleteUserInterestResponse{}, err
	}

	return &pb.DeleteUserInterestResponse{
		Status: 201,
	}, nil
}
