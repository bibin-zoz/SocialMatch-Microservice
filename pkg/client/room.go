package client

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/config"
	"github.com/bibin-zoz/api-gateway/pkg/pb/room"
	pb "github.com/bibin-zoz/api-gateway/pkg/pb/room"
	"github.com/bibin-zoz/api-gateway/pkg/utils/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RoomClient struct {
	Client     pb.RoomServiceClient
	RoomClient room.RoomServiceClient
}

func NewRoomServiceClient(cfg config.Config) interfaces.RoomClient {
	grpcConnection, err := grpc.NewClient(cfg.RoomSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect", err)
		return nil
	}

	grpcClient := pb.NewRoomServiceClient(grpcConnection)
	roomGrpcConnection, err := grpc.NewClient(cfg.RoomSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect to room service", err)
		return nil
	}
	roomGrpcClient := room.NewRoomServiceClient(roomGrpcConnection)
	fmt.Println("grpc", grpcClient)
	return &RoomClient{
		Client:     grpcClient,
		RoomClient: roomGrpcClient,
	}

}

func (c *RoomClient) CreateRoom(roomData models.RoomData) (models.Room, error) {
	res, err := c.RoomClient.CreateRoom(context.Background(), &room.RoomCreateRequest{
		RoomName:    roomData.RoomName,
		Description: roomData.Description,
		MaxPersons:  roomData.MaxPersons,
		Status:      roomData.Status,
		Preferences: roomData.Preferences,
	})
	if err != nil {
		return models.Room{}, err
	}
	return models.Room{
		ID:          res.Id,
		RoomName:    res.RoomName,
		Description: res.Description,
		MaxPersons:  res.MaxPersons,
		Status:      res.Status,
		// Preferences: res.P,
	}, nil
}

func (c *RoomClient) EditRoom(roomData models.RoomData) (models.Room, error) {
	res, err := c.RoomClient.EditRoom(context.Background(), &room.RoomEditRequest{
		RoomId:      roomData.ID,
		RoomName:    roomData.RoomName,
		Description: roomData.Description,
		MaxPersons:  roomData.MaxPersons,
		Status:      roomData.Status,
		Preferences: roomData.Preferences,
	})
	if err != nil {
		return models.Room{}, err
	}
	return models.Room{
		ID:          res.Id,
		RoomName:    res.RoomName,
		Description: res.Description,
		MaxPersons:  res.MaxPersons,
		Status:      res.Status,
		// Preferences: res.Preferences,
	}, nil
}

func (c *RoomClient) ChangeRoomStatus(roomID uint32, status string) (models.Room, error) {
	res, err := c.RoomClient.ChangeStatus(context.Background(), &room.ChangeStatusRequest{
		RoomId: roomID,
		Status: status,
	})
	if err != nil {
		return models.Room{}, err
	}
	return models.Room{
		ID:          res.Id,
		RoomName:    res.RoomName,
		Description: res.Description,
		MaxPersons:  res.MaxPersons,
		Status:      res.Status,
		// Preferences: res.Preferences,
	}, nil
}

func (c *RoomClient) AddMembersToRoom(roomID uint32, userIds []uint32) (models.Room, error) {
	res, err := c.RoomClient.AddMembers(context.Background(), &room.AddMembersRequest{
		RoomId:  roomID,
		UserIds: userIds,
	})
	if err != nil {
		return models.Room{}, err
	}
	return models.Room{
		ID:          res.Id,
		RoomName:    res.RoomName,
		Description: res.Description,
		MaxPersons:  res.MaxPersons,
		Status:      res.Status,
		// Preferences: res.Preferences,
	}, nil

}

func (c *RoomClient) GetRoomJoinRequests(roomID uint32) ([]models.RoomJoinRequest, error) {
	res, err := c.RoomClient.SeeRoomJoinRequests(context.Background(), &room.RoomJoinRequestsRequest{
		RoomId: roomID,
	})
	if err != nil {
		return nil, err
	}
	joinRequests := make([]models.RoomJoinRequest, len(res.RoomJoinRequests))
	for i, req := range res.RoomJoinRequests {
		joinRequests[i] = models.RoomJoinRequest{
			UserID:   req.UserId,
			Username: req.Username,
		}
	}
	return joinRequests, nil
}
