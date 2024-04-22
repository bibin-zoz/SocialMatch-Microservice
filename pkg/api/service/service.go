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
			Id:           (u.Id),
			InterestName: u.InterestName,
		}
		pbInterest = append(pbInterest, pbRes)
	}

	return &pb.GetInterestsResponse{
		Status:    201,
		Interests: pbInterest,
	}, nil
}
