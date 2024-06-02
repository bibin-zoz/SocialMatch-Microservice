// client/interest.go

package client

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/config"
	interest "github.com/bibin-zoz/social-match-userauth-svc/pkg/pb/interest"
	"google.golang.org/grpc"
)

// InterestClient is a client for interacting with the InterestService.
type InterestClient struct {
	conn    *grpc.ClientConn
	service interest.InterestServiceClient
}

// NewInterestClient creates a new instance of InterestClient.
func NewInterestClient(cfg config.Config) *InterestClient {
	// Print debug information
	fmt.Println("Dialing gRPC server at:", cfg.AdminSvcUrl)

	// Dial the gRPC server
	conn, err := grpc.Dial(cfg.AdminSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to dial gRPC server:", err)
		return nil
	}

	// Create a client for the InterestService
	service := interest.NewInterestServiceClient(conn)

	return &InterestClient{
		conn:    conn,
		service: service,
	}
}

func (c *InterestClient) CheckInterest(interestID string) (bool, error) {
	request := &interest.CheckInterestRequest{
		InterestId: interestID,
	}
	response, err := c.service.CheckInterest(context.Background(), request)
	if err != nil {
		return false, err
	}
	return response.Exists, nil
}
