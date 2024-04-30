// service/interestservice.go

package service

import (
	"context"

	pb "github.com/bibin-zoz/social-match-admin-svc/pkg/pb/interest"
)

type InterestService struct {
	useCase interfaces.InterestUseCase // Assuming you have an interface for the InterestUseCase
	pb.UnimplementedInterestServer
}

func NewInterestService(useCase interfaces.InterestUseCase) *InterestService {
	return &InterestService{
		useCase: useCase,
	}
}

func (is *InterestService) CheckInterest(ctx context.Context, req *pb.CheckInterestRequest) (*pb.CheckInterestResponse, error) {
	// Perform the logic to check if the interest exists using the interest use case
	exists, err := is.useCase.CheckInterest(ctx, req.InterestId)
	if err != nil {
		return nil, err
	}
	return &pb.CheckInterestResponse{Exists: exists}, nil
}
