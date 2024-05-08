package helper

import (
	"strconv"
)

func generateRoomID(user1, user2 int64) string {
	// Sort user IDs to ensure consistency
	if user1 > user2 {
		user1, user2 = user2, user1
	}
	return strconv.FormatInt(user1, 10) + "-" + strconv.FormatInt(user2, 10)
}

// func broadcastMessages(roomID string) {
// 	room := rooms[roomID]
// 	for msg := range room.Ch {
// 		for _, conn := range room.Connections {
// 			if err := conn.WriteJSON(msg); err != nil {
// 				log.Printf("Error sending message to connection: %v", err)
// 			}
// 		}
// 	}
// }
// func getPreviousMessages(userID, receiverID int64) ([]*Message, error) {
// 	var messages []*Message
// 	if err := db.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", userID, receiverID, receiverID, userID).Order("time").Find(&messages).Error; err != nil {
// 		return nil, err
// 	}
// 	return messages, nil
// }
