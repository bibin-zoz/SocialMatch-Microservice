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
	grpcConnection, err := grpc.Dial(cfg.RoomSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect", err)
		return nil
	}

	grpcClient := pb.NewRoomServiceClient(grpcConnection)
	roomGrpcConnection, err := grpc.Dial(cfg.RoomSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

func (c *RoomClient) CreateRoom(roomData models.RoomData, userid int) (models.Room, error) {
	res, err := c.RoomClient.CreateRoom(context.Background(), &room.RoomCreateRequest{
		RoomName:    roomData.RoomName,
		Description: roomData.Description,
		MaxPersons:  roomData.MaxPersons,
		Status:      roomData.Status,
		Preferences: roomData.Preferences,
		CreatorID:   int32(userid),
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

func (c *RoomClient) GetAllRooms() ([]models.Room, error) {
	res, err := c.Client.GetAllRooms(context.Background(), &pb.GetAllRoomsRequest{})
	if err != nil {
		return nil, err
	}

	var rooms []models.Room
	for _, pbRoom := range res.Rooms {
		room := models.Room{
			ID:          pbRoom.Id,
			RoomName:    pbRoom.RoomName,
			Description: pbRoom.Description,
			MaxPersons:  pbRoom.MaxPersons,
			Status:      pbRoom.Status,
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (c *RoomClient) GetRoomMembers(roomID uint32) ([]models.RoomMember, error) {
	res, err := c.RoomClient.GetGroupMembers(context.Background(), &room.GetGroupMembersRequest{RoomId: roomID})
	if err != nil {
		return nil, err
	}

	var members []models.RoomMember
	for _, pbMember := range res.Members {
		member := models.RoomMember{
			UserID:   pbMember.UserId,
			Username: pbMember.Username,
		}
		members = append(members, member)
	}

	return members, nil
}
func (c *RoomClient) SendMessage(message models.Message) (models.Message, error) {
	req := &room.SendMessageRequest{
		RoomId:   uint32(message.RoomID),
		SenderId: uint32(message.UserID),
		Content:  message.Content,
		Media:    message.Media,
	}
	res, err := c.RoomClient.SendMessage(context.Background(), req)
	if err != nil {
		return models.Message{}, err
	}
	return models.Message{
		ID:        uint(res.MessageId),
		RoomID:    uint(res.RoomId),
		UserID:    uint(res.SenderId),
		Content:   res.Content,
		CreatedAt: res.Timestamp.AsTime(), // Convert Protobuf timestamp to Go time
	}, nil
}

func (c *RoomClient) ReadMessages(roomID uint32) ([]models.Message, error) {
	req := &room.ReadMessagesRequest{
		RoomId: roomID,
	}
	res, err := c.RoomClient.ReadMessages(context.Background(), req)
	if err != nil {
		return nil, err
	}
	messages := make([]models.Message, len(res.Messages))
	for i, pbMsg := range res.Messages {
		messages[i] = models.Message{
			ID:        uint(pbMsg.MessageId),
			RoomID:    uint(pbMsg.RoomId),
			UserID:    uint(pbMsg.SenderId),
			Content:   pbMsg.Content,
			CreatedAt: pbMsg.Timestamp.AsTime(), // Convert Protobuf timestamp to Go time
		}
	}
	return messages, nil
}
