// server/interest.go

package service

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/social-match-admin-svc/pkg/pb/interest"
	pb "github.com/bibin-zoz/social-match-admin-svc/pkg/pb/interest"
	interfaces "github.com/bibin-zoz/social-match-admin-svc/pkg/usecase/interface"
)

// InterestServiceServer is the server for the InterestService.
type InterestServiceServer struct {
	interestUsecase interfaces.InterestUseCase
	pb.UnimplementedInterestServiceServer
}

// NewInterestServiceServer creates a new instance of InterestServiceServer.
func NewInterestServiceServer(usecase interfaces.InterestUseCase) pb.InterestServiceServer {
	return &InterestServiceServer{
		interestUsecase: usecase,
	}
}

func (s *InterestServiceServer) CheckInterest(ctx context.Context, req *interest.CheckInterestRequest) (*interest.CheckInterestResponse, error) {
	fmt.Println("entered")
	exists, _ := s.interestUsecase.CheckInterest(req.InterestId)
	fmt.Println(exists)
	return &interest.CheckInterestResponse{
		Exists: exists,
	}, nil
}
