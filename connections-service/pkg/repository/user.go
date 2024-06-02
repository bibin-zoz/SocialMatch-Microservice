package repository

import (
	"errors"

	"github.com/bibin-zoz/social-match-connections-svc/pkg/domain"
	interfaces "github.com/bibin-zoz/social-match-connections-svc/pkg/repository/interface"
	"github.com/bibin-zoz/social-match-connections-svc/pkg/utils/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type ConnectionRepository struct {
	DB          *gorm.DB
	MongoClient *mongo.Client
}

func NewConnectionRepository(DB *gorm.DB, mongoClient *mongo.Client) interfaces.ConnectionRepository {
	return &ConnectionRepository{
		DB:          DB,
		MongoClient: mongoClient,
	}

}

//	func (ur *ConnectionRepository) FollowUser(senderID uint64, userID uint64) ([]string, error) {
//		var followReq domain.FollowUSer
//		fmt.Println("id", userID)
//		err := ur.DB.Model(&domain.UserInterest{}).Where("user_id = ?", userID).Pluck("interest_id", &interests).Error
//		if err != nil {
//			return nil, err
//		}
//		return interests, nil
//	}
func (ur *ConnectionRepository) CheckConnectionRequestExist(senderID, userID uint) (bool, error) {
	var connection domain.Connections
	result := ur.DB.Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", senderID, userID, userID, senderID).First(&connection)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// AddConnection adds a connection between senderID and userID
func (ur *ConnectionRepository) AddConnection(senderID, userID uint) error {
	connection := domain.Connections{
		UserID:   senderID,
		FriendID: userID,
		Status:   "friends",
	}
	result := ur.DB.Where("user_id=? and friend_id=?", senderID, userID).Or("user_id=? and friend_id=?", userID, senderID).Updates(&connection)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (ur *ConnectionRepository) BlockConnection(senderID, userID uint) error {
	connection := domain.Connections{
		UserID:   senderID,
		FriendID: userID,
		Status:   "blocked",
	}
	result := ur.DB.Where("user_id=? and friend_id=?", senderID, userID).Or("user_id=? and friend_id=?", userID, senderID).Updates(&connection)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CheckFriends checks if senderID and userID are already friends
func (ur *ConnectionRepository) CheckFriends(senderID, userID uint) (bool, error) {
	var connection domain.Connections
	result := ur.DB.Where("(user_id = ? AND friend_id = ?) AND status = 'friends'", senderID, userID).First(&connection)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// AddConnectionRequest adds a connection request between senderID and userID
func (ur *ConnectionRepository) AddConnectionRequest(senderID, userID uint) error {
	connection := domain.Connections{
		UserID:   senderID,
		FriendID: userID,
		Status:   "pending",
	}
	result := ur.DB.Create(&connection)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (ur *ConnectionRepository) GetConnections(userID uint) ([]models.UserDetails, error) {
	// Query the database to fetch connections for the specified user ID
	var connections []models.UserDetails
	err := ur.DB.Table("connections").
		Where("connections.user_id = ? AND connections.status = 'friends'", userID).Or("connections.friend_id = ? AND connections.status = 'friends'", userID).
		Find(&connections).Error
	if err != nil {
		return nil, err
	}
	return connections, nil
}

func (ur *ConnectionRepository) CheckUserConnection(userID int64, recvID int64) (bool, error) {
	var count int64
	err := ur.DB.Table("connections").
		Where("(user_id = ? AND friend_id = ? OR user_id = ? AND friend_id = ?) AND status = 'friends'",
			userID, recvID, recvID, userID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
