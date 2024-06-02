package client

import (
	"context"
	"fmt"
	"log"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/config"
	pb "github.com/bibin-zoz/social-match-userauth-svc/pkg/pb/chat"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
	"google.golang.org/grpc"
)

// ChatClient is a client for interacting with the ChatService.
type ChatClient struct {
	conn    *grpc.ClientConn
	service pb.ChatServiceClient
}

// NewChatClient creates a new instance of ChatClient.
func NewChatClient(cfg config.Config) *ChatClient {

	// Print debug information
	fmt.Println("Dialing gRPC server at:", cfg.ChatSvcUrl)

	// Dial the gRPC server
	conn, err := grpc.Dial(cfg.ChatSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial gRPC server: %v", err)
		return nil
	}

	// Create a client for the ChatService
	service := pb.NewChatServiceClient(conn)

	return &ChatClient{
		conn:    conn,
		service: service,
	}
}

// SendMessage calls the SendMessage method on the ChatService.
func (c *ChatClient) SendMessage(senderID, receiverID uint32, content string, media []models.Media) (*pb.SendMessageResponce, error) {
	request := &pb.SendMessageRequest{
		SenterId:   senderID,
		ReceiverId: receiverID,
		Content:    content,
		Media:      nil,
	}
	response, err := c.service.SendMessage(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// ReadMessages calls the ReadMessages method on the ChatService.
func (c *ChatClient) ReadMessages(senderID, receiverID uint32) ([]*pb.Message, error) {
	request := &pb.ReadMessagesRequest{
		SenterId:   senderID,
		ReceiverId: receiverID,
	}
	response, err := c.service.ReadMessages(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response.Messages, nil
}
