package client

import (
	pb "github.com/bibin-zoz/social-match-userauth-svc/pkg/pb/chat"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
)

type ChatClientInterface interface {
	SendMessage(senderID, receiverID uint32, content string, media []models.Media) (*pb.SendMessageResponce, error)
	ReadMessages(senderID, receiverID uint32) ([]*pb.Message, error)
}
