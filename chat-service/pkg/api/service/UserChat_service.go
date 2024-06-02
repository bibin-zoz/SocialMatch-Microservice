package service

import (
	"context"
	"time"

	pb "github.com/bibin-zoz/social-match-chat-svc/pkg/pb/chat"
	"github.com/bibin-zoz/social-match-chat-svc/pkg/utils/models"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"

	interfaces "github.com/bibin-zoz/social-match-chat-svc/pkg/usecase/interface"
)

type ChatServer struct {
	chatUseCase interfaces.ChatUseCase
	pb.UnimplementedChatServiceServer
}

func NewChatServer(useCase interfaces.ChatUseCase) pb.ChatServiceServer {

	return &ChatServer{
		chatUseCase: useCase,
	}

}
func (a *ChatServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponce, error) {
	media := make([]models.Media, len(req.Media))
	for i, m := range req.Media {
		media[i] = models.Media{
			Filename: m.Filename,
		}
	}

	message := &models.UserMessage{
		SenderID:   uint(req.SenterId),
		RecipentID: uint(req.ReceiverId),
		Content:    req.Content,
		CreatedAt:  time.Now(),
		Media:      media,
	}
	err := a.chatUseCase.SendMessage(message)
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponce{
		MessageId:  123,
		SenterId:   req.SenterId,
		ReceiverId: req.ReceiverId,
		Content:    req.Content,
		Timestamp:  &timestamp.Timestamp{},
		Media:      req.Media,
	}, nil
}
func (a *ChatServer) ReadMessages(ctx context.Context, req *pb.ReadMessagesRequest) (*pb.ReadMessagesResponse, error) {
	// Call the use case method to get messages for the specified room
	messages, err := a.chatUseCase.GetMessages(uint64(req.ReceiverId), uint64(req.SenterId))
	if err != nil {
		return nil, err
	}

	// Convert the returned messages to protobuf format and return
	var pbMessages []*pb.Message
	for _, message := range messages {
		timestamp := timestamppb.New(message.CreatedAt)
		pbMessage := &pb.Message{
			MessageId:    uint32(message.ID),
			ReceiverIdId: uint32(message.RecipentID),
			SenderId:     uint32(message.SenderID),
			Content:      message.Content,
			Timestamp:    timestamp,
		}
		pbMessages = append(pbMessages, pbMessage)
	}

	return &pb.ReadMessagesResponse{
		Messages: pbMessages,
	}, nil
}
