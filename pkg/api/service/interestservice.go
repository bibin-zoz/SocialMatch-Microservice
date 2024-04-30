// server/interest.go

package service

import (
	"context"
)

// InterestServiceServer is the server for the InterestService.
type InterestServiceServer struct {
	// You can add any dependencies or repository interfaces here
}

// NewInterestServiceServer creates a new instance of InterestServiceServer.
func NewInterestServiceServer() *InterestServiceServer {
	// If you need to initialize anything, you can do it here
	return &InterestServiceServer{}
}

// CheckInterest checks if an interest ID exists.
func (s *InterestServiceServer) CheckInterest(ctx context.Context, req *interest.CheckInterestRequest) (*interest.CheckInterestResponse, error) {
	// Perform logic to check if the interest ID exists in the database or elsewhere
	// For simplicity, let's assume the interest ID exists if it's not empty
	exists := req.InterestId != ""

	// Return the response
	return &interest.CheckInterestResponse{
		Exists: exists,
	}, nil
}
