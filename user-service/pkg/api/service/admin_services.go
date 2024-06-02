package service

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/pb"
)

func (a *UserServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := a.userUseCase.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.Users
	for _, u := range users {
		pbUser := &pb.Users{
			Id:        uint32(u.ID),
			Firstname: u.Firstname,
			Lastname:  u.Lastname,
			Email:     u.Email,
			Password:  u.Password,
			Phone:     u.Phone,
			Blocked:   u.Blocked,
			Username:  u.Username,
			GenderId:  int32(u.GenderID),
			Age:       int32(u.Age),
		}
		pbUsers = append(pbUsers, pbUser)
	}

	return &pb.GetUsersResponse{
		Status: 201,
		Users:  pbUsers,
	}, nil
}

func (a *UserServer) UpdateUserStatus(ctx context.Context, req *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusResponse, error) {
	fmt.Println("hiiUpdateUserStatus")
	_, err := a.userUseCase.UpdateUserStatus(int(req.Userid))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserStatusResponse{
		Status: 201,
	}, nil
}
