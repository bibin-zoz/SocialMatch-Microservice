package service

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/social-match-room-svc/pkg/domain"
	pb "github.com/bibin-zoz/social-match-room-svc/pkg/pb/room"
	interfaces "github.com/bibin-zoz/social-match-room-svc/pkg/usecase/interface"
	"github.com/bibin-zoz/social-match-room-svc/pkg/utils/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RoomServer struct {
	roomUseCase interfaces.RoomUseCase
	pb.UnimplementedRoomServiceServer
}

func NewRoomServer(useCase interfaces.RoomUseCase) pb.RoomServiceServer {
	return &RoomServer{
		roomUseCase: useCase,
	}
}

func (s *RoomServer) CreateRoom(ctx context.Context, req *pb.RoomCreateRequest) (*pb.Room, error) {
	newRoom := domain.Room{
		RoomName:    req.RoomName,
		Description: req.Description,
		MaxPersons:  uint(req.MaxPersons),
		Status:      req.Status,
		CreatorID:   int(req.CreatorID),
	}
	fmt.Println("reqqqqq", req.CreatorID)

	room, err := s.roomUseCase.CreateRoom(newRoom)
	if err != nil {
		return nil, err
	}

	pbRoom := &pb.Room{
		Id:          uint32(room.ID),
		RoomName:    room.RoomName,
		Description: room.Description,
		MaxPersons:  uint32(room.MaxPersons),
		Status:      room.Status,
	}

	return pbRoom, nil
}

func (s *RoomServer) EditRoom(ctx context.Context, req *pb.RoomEditRequest) (*pb.Room, error) {
	roomEditDetails := models.RoomEdit{
		RoomID:      uint(req.RoomId),
		RoomName:    req.RoomName,
		Description: req.Description,
		MaxPersons:  uint(req.MaxPersons),
		Status:      req.Status,
	}

	room, err := s.roomUseCase.EditRoom(roomEditDetails)
	if err != nil {
		return nil, err
	}

	pbRoom := &pb.Room{
		Id:          uint32(room.ID),
		RoomName:    room.RoomName,
		Description: room.Description,
		MaxPersons:  uint32(room.MaxPersons),
		Status:      room.Status,
	}

	return pbRoom, nil
}

func (s *RoomServer) ChangeStatus(ctx context.Context, req *pb.ChangeStatusRequest) (*pb.Room, error) {
	room, err := s.roomUseCase.ChangeStatus(uint(req.RoomId), req.Status)
	if err != nil {
		return nil, err
	}

	pbRoom := &pb.Room{
		Id:     uint32(room.ID),
		Status: room.Status,
	}

	return pbRoom, nil
}

func (s *RoomServer) AddMembers(ctx context.Context, req *pb.AddMembersRequest) (*pb.Room, error) {
	room, err := s.roomUseCase.AddMembers(uint(req.RoomId), req.UserIds)
	if err != nil {
		return nil, err
	}

	pbRoom := &pb.Room{
		Id:          uint32(room.ID),
		RoomName:    room.RoomName,
		Description: room.Description,
		MaxPersons:  uint32(room.MaxPersons),
		Status:      room.Status,
	}

	return pbRoom, nil
}
func (s *RoomServer) GetAllRooms(ctx context.Context, req *pb.GetAllRoomsRequest) (*pb.GetAllRoomsResponse, error) {
	rooms, err := s.roomUseCase.GetAllRooms() // Call the use case method
	if err != nil {
		return nil, err
	}

	var pbRooms []*pb.Room
	for _, room := range rooms {
		pbRoom := &pb.Room{
			Id:          uint32(room.ID),
			RoomName:    room.RoomName,
			Description: room.Description,
			MaxPersons:  uint32(room.MaxPersons),
			Status:      room.Status,
			// Preferences: room.Preferences, // If needed
		}
		pbRooms = append(pbRooms, pbRoom)
	}

	return &pb.GetAllRoomsResponse{
		Rooms: pbRooms,
	}, nil
}
func (s *RoomServer) SeeRoomJoinRequests(ctx context.Context, req *pb.RoomJoinRequestsRequest) (*pb.RoomJoinRequestsResponse, error) {
	roomJoinRequests, err := s.roomUseCase.SeeRoomJoinRequests(req.RoomId)
	if err != nil {
		return nil, err
	}

	var pbRoomJoinRequests []*pb.RoomJoinRequest
	for _, req := range roomJoinRequests {
		pbRoomJoinRequests = append(pbRoomJoinRequests, &pb.RoomJoinRequest{
			UserId:   req.UserID,
			Username: req.Username,
		})
	}

	return &pb.RoomJoinRequestsResponse{
		RoomJoinRequests: pbRoomJoinRequests,
	}, nil
}

func (s *RoomServer) GetGroupMembers(ctx context.Context, req *pb.GetGroupMembersRequest) (*pb.GetGroupMembersResponse, error) {
	roomMembers, err := s.roomUseCase.GetRoomMembers(req.RoomId)
	if err != nil {
		return nil, err
	}

	var pbRoomMembers []*pb.RoomJoinRequest
	for _, member := range roomMembers {
		pbRoomMember := &pb.RoomJoinRequest{
			UserId:   uint32(member.UserID),
			Username: member.Username,
		}
		pbRoomMembers = append(pbRoomMembers, pbRoomMember)
	}

	return &pb.GetGroupMembersResponse{
		Members: pbRoomMembers,
	}, nil
}

func (s *RoomServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.Message, error) {

	// Convert the request to a domain message
	var mediaList []domain.Media
	for _, m := range req.Media {
		mediaList = append(mediaList, domain.Media{
			ID:       uint(m.Id),
			Filename: m.Filename,
			// Add other field conversions as needed
		})
	}
	newMessage := domain.Message{
		RoomID:  uint(req.RoomId),
		UserID:  uint(req.SenderId),
		Content: req.Content,
		Media:   mediaList,
	}

	// Call the use case method to send the message
	message, err := s.roomUseCase.SendMessage(newMessage)
	if err != nil {
		return nil, err
	}
	timestamp := timestamppb.New(message.CreatedAt)
	pbMessage := &pb.Message{
		MessageId: uint32(message.ID),
		RoomId:    uint32(message.RoomID),
		SenderId:  uint32(message.UserID),
		Content:   message.Content,
		Timestamp: timestamp,
		Media:     req.Media,
	}

	return pbMessage, nil
}

func (s *RoomServer) ReadMessages(ctx context.Context, req *pb.ReadMessagesRequest) (*pb.ReadMessagesResponse, error) {
	// Call the use case method to get messages for the specified room
	messages, err := s.roomUseCase.GetMessages(uint(req.RoomId))
	if err != nil {
		return nil, err
	}

	// Convert the returned messages to protobuf format and return
	var pbMessages []*pb.Message
	for _, message := range messages {
		timestamp := timestamppb.New(message.CreatedAt)
		pbMessage := &pb.Message{
			MessageId: uint32(message.ID),
			RoomId:    uint32(message.RoomID),
			SenderId:  uint32(message.UserID),
			Content:   message.Content,
			Timestamp: timestamp,
		}
		pbMessages = append(pbMessages, pbMessage)
	}

	return &pb.ReadMessagesResponse{
		Messages: pbMessages,
	}, nil
}
