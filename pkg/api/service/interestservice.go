// server/interest.go

package service

import (
	"context"

	pb "github.com/bibin-zoz/social-match-admin-svc/pkg/pb/admin"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/pb/interest"
	interfaces "github.com/bibin-zoz/social-match-admin-svc/pkg/usecase/interface"
)

// InterestServiceServer is the server for the InterestService.
type InterestServiceServer struct {
	interestUsecase interfaces.InterestUseCase
	pb.UnimplementedAdminServer
}

// NewInterestServiceServer creates a new instance of InterestServiceServer.
func NewInterestServiceServer(usecase interfaces.InterestUseCase) *InterestServiceServer {
	return &InterestServiceServer{
		interestUsecase: usecase,
	}
}

func (s *InterestServiceServer) CheckInterest(ctx context.Context, req *interest.CheckInterestRequest) (*interest.CheckInterestResponse, error) {
	exists, _ := s.interestUsecase.CheckInterest(req.InterestId)

	return &interest.CheckInterestResponse{
		Exists: exists,
	}, nil
}
