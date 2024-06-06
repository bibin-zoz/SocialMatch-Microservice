package client

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/config"
	pb "github.com/bibin-zoz/social-match-userauth-svc/pkg/pb/connections"
	"google.golang.org/grpc"
)

// ConnectionsClient is a client for interacting with the ConnectionsService.
type ConnectionsClient struct {
	conn    *grpc.ClientConn
	service pb.ConnectionsClient
}

// NewConnectionsClient creates a new instance of ConnectionsClient.
func NewConnectionsClient(cfg config.Config) *ConnectionsClient {
	// Print debug information
	fmt.Println("Dialing gRPC server at:", cfg.ConnectionSvcUrl)

	// Dial the gRPC server
	conn, err := grpc.Dial(cfg.ConnectionSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println(fmt.Errorf("failed to dial gRPC server: %w", err))
		return nil

	}

	// Create a client for the ConnectionsService
	service := pb.NewConnectionsClient(conn)

	return &ConnectionsClient{
		conn:    conn,
		service: service,
	}
}

// FollowUser calls the FollowUser method on the ConnectionsService.
func (c *ConnectionsClient) FollowUser(userID, senderID int64) (int64, error) {
	request := &pb.FollowUserRequest{
		Userid:   userID,
		Senderid: senderID,
	}
	response, err := c.service.FollowUser(context.Background(), request)
	if err != nil {
		return 0, err
	}
	return response.Status, nil
}

// BlockUser calls the BlockUser method on the ConnectionsService.
func (c *ConnectionsClient) BlockUserConnection(userID, senderID int64) (int64, error) {
	request := &pb.BlockConnectionRequest{
		UserId:   uint64(userID),
		SenderId: uint64(senderID),
	}
	response, err := c.service.BlockConnection(context.Background(), request)
	if err != nil {
		return 0, err
	}
	return int64(response.Status), nil
}

// GetConnections calls the GetConnections method on the ConnectionsService.
func (c *ConnectionsClient) GetConnections(userID uint64) ([]*pb.UserDetails, error) {
	fmt.Println("user connect")
	request := &pb.GetConnectionsRequest{
		UserId: userID,
	}
	response, err := c.service.GetConnections(context.Background(), request)
	if err != nil {
		return nil, err
	}
	fmt.Println("res", response)
	return response.UserDetails, nil
}
