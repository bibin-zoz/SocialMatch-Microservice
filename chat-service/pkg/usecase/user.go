package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bibin-zoz/social-match-chat-svc/pkg/config"
	"github.com/bibin-zoz/social-match-chat-svc/pkg/domain"
	interfaces "github.com/bibin-zoz/social-match-chat-svc/pkg/repository/interface"
	services "github.com/bibin-zoz/social-match-chat-svc/pkg/usecase/interface"
	"github.com/bibin-zoz/social-match-chat-svc/pkg/utils/models"
	"github.com/segmentio/kafka-go"
)

type chatUseCase struct {
	chatRepository interfaces.ChatRepository
	Config         config.Config
}

func NewChatUseCase(repository interfaces.ChatRepository, config config.Config) services.ChatUseCase {
	return &chatUseCase{
		chatRepository: repository,
		Config:         config,
	}
}
func (ch *chatUseCase) SendMessage(message *models.UserMessage) error {
	fmt.Println("usecase")
	// Assuming you have a repository method to save the message
	usermsg := &domain.UserMessage{
		SenderID:   message.SenderID,
		RecipentID: message.RecipentID,
		Content:    message.Content,
		CreatedAt:  message.CreatedAt,
	}
	msgID, err := ch.chatRepository.SaveMessage(usermsg)
	if err != nil {
		return errors.New("failed to save message")
	}
	if message.Media != nil {

		for _, m := range message.Media {
			media := &domain.Media{
				Message_id: int(msgID),
				Filename:   m.Filename,
			}
			fmt.Println("hi2")
			err := ch.chatRepository.SaveMedia(media)
			if err != nil {
				return errors.New("failed to save message")
			}
		}

	}
	fmt.Println("messaged saved successfully")
	return nil
}
func (ch *chatUseCase) ConsumeAndProcessMessages() {
	config := kafka.ReaderConfig{
		Brokers: []string{ch.Config.KafkaBrokers},
		GroupID: "1",
		Topic:   "chat",
	}

	consumer := kafka.NewReader(config)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signals)

	// Handle signals
	go func() {
		sig := <-signals
		fmt.Printf("Received signal: %v\n", sig)
		fmt.Println("Shutting down...")

		// Close Kafka consumer
		if err := consumer.Close(); err != nil {
			fmt.Printf("Error closing Kafka consumer: %v\n", err)
		}

		os.Exit(0)
	}()

	// Continuously consume messages
	for {
		msg, err := consumer.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}

		var userMessage models.UserMessage
		if err := json.Unmarshal(msg.Value, &userMessage); err != nil {
			fmt.Printf("Failed to deserialize message: %v\n", err)
			continue
		}

		usermsg := &models.UserMessage{
			SenderID:   userMessage.SenderID,
			RecipentID: userMessage.RecipentID,
			Content:    userMessage.Content,
			CreatedAt:  userMessage.CreatedAt,
		}

		if err := ch.SendMessage(usermsg); err != nil {
			fmt.Printf("Failed to process message: %v\n", err)
		}
		fmt.Println("usermsg", userMessage)
	}
}

func (ch *chatUseCase) GetMessages(receiverID, senderID uint64) ([]models.UserMessage, error) {
	messages, err := ch.chatRepository.GetMessages(receiverID, senderID)
	if err != nil {
		return nil, errors.New("error retrieving messages")
	}
	return messages, nil
}
