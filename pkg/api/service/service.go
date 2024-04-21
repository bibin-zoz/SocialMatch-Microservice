package service

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/social-match-admin-svc/pkg/pb"
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

func (ad *AdminServer) AdminLogin(ctx context.Context, Req *pb.AdminLoginInRequest) (*pb.AdminLoginResponse, error) {
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
