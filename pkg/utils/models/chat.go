package models

import "github.com/gorilla/websocket"

type WebrtcMessage struct {
	SenderID   int64  `json:"sender_id"`
	ReceiverID int64  `json:"receiver_id"`
	Message    string `json:"message"`
	// Time       time.Time `json:"time"`
}

type User struct {
	UserID int64  `json:"user_id" gorm:"primary_key"`
	Name   string `json:"name"`
}

type WebrtcRoom struct {
	User1       int64
	User2       int64
	Connections []*websocket.Conn
	Ch          chan *WebrtcMessage
}
