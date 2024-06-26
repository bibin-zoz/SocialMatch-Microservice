package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/domain"
	interfaces "github.com/bibin-zoz/social-match-userauth-svc/pkg/repository/interface"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type userRepository struct {
	DB      *gorm.DB
	MongoDB *mongo.Database
}

func NewUserRepository(DB *gorm.DB, MongoDB *mongo.Database) interfaces.UserRepository {
	return &userRepository{
		DB:      DB,
		MongoDB: MongoDB,
	}

}
func (ur *userRepository) SaveProfilePhoto(userID uint32, imageURLs []string) error {
	fmt.Println("repo entered")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := bson.M{
		"userid":         userID,
		"profile_photos": imageURLs,
	}

	_, err := ur.MongoDB.Collection("userimages").InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to save profile photos: %v", err)
	}

	return nil
}

func (ur *userRepository) CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := ur.DB.Where(&domain.User{Email: email}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.User{}, res.Error
	}
	return &user, nil
}
func (ur *userRepository) CheckUserExistsByPhone(phone string) (*domain.User, error) {
	var user domain.User
	res := ur.DB.Where(&domain.User{Phone: phone}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.User{}, res.Error
	}
	return &user, nil
}
func (ur *userRepository) UserSignUp(user models.UserSignup) (models.UserDetails, error) {
	var signupDetail models.UserDetails
	err := ur.DB.Raw(`
		INSERT INTO users(firstname, lastname, email, password, phone)
		VALUES(?, ?, ?, ?, ?)
		RETURNING id, firstname, lastname, email, phone
	`, user.FirstName, user.LastName, user.Email, user.Password, user.Number).
		Scan(&signupDetail).Error
	fmt.Println("pass", user.Password)
	if err != nil {
		return models.UserDetails{}, err
	}
	return signupDetail, nil
}
func (ur *userRepository) FindUserByEmail(user models.UserLogin) (models.UserDetail, error) {
	var userDetails models.UserDetail
	err := ur.DB.Raw("SELECT * FROM users WHERE email=?", user.Email).Scan(&userDetails).Error
	if err != nil {
		return models.UserDetail{}, errors.New("error checking user details")
	}
	return userDetails, nil
}

func (ur *userRepository) UserEditDetails(user models.UserSignup) (models.UserDetails, error) {
	var userDetails models.UserDetails
	err := ur.DB.Raw(`
		UPDATE users 
		SET firstname = ?, lastname = ?, password = ?
		WHERE EMAIL = ?
		RETURNING id, firstname, lastname, email, phone
	`, user.FirstName, user.LastName, user.Password, user.Email).
		Scan(&userDetails).Error

	if err != nil {
		return models.UserDetails{}, err
	}
	return userDetails, nil
}
func (ur *userRepository) FetchAllUsers() ([]domain.User, error) {
	var users []domain.User
	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	fmt.Println("users", users)
	return users, nil
}

func (ur *userRepository) GetUserByID(id int) (domain.User, error) {
	var user domain.User
	if err := ur.DB.First(&user, id).Error; err != nil {
		return domain.User{}, err
	}
	fmt.Println("User:", user)
	return user, nil
}

func (ur *userRepository) BlockUser(userID uint) error {
	return ur.DB.Model(&domain.User{}).Where("id = ?", userID).Update("blocked", true).Error
}

func (ur *userRepository) UnblockUser(userID uint) error {
	return ur.DB.Model(&domain.User{}).Where("id = ?", userID).Update("blocked", false).Error
}
func (ur *userRepository) CheckUserInterest(userID uint64, interestName string) (bool, error) {
	var count int64
	err := ur.DB.Model(&domain.UserInterest{}).Where("user_id = ? AND interest_name = ?", userID, interestName).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ur *userRepository) CheckUserInterestByID(userID uint64, interestID uint64) (bool, error) {
	var count int64
	err := ur.DB.Model(&domain.UserInterest{}).Where("user_id = ? AND interest_id = ?", userID, interestID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ur *userRepository) AddUserInterest(userID uint64, interestID int) error {
	interest := domain.UserInterest{
		UserID:     int(userID),
		InterestID: interestID,
	}
	return ur.DB.Create(&interest).Error
}

func (ur *userRepository) EditUserInterest(userID uint64, interestID uint64, newInterestName string) error {
	return ur.DB.Model(&domain.UserInterest{}).Where("user_id = ? AND id = ?", userID, interestID).Update("interest_name", newInterestName).Error
}

func (ur *userRepository) DeleteUserInterest(userID uint64, interestID uint64) error {
	return ur.DB.Delete(&domain.UserInterest{}, "user_id = ? AND id = ?", userID, interestID).Error
}
func (ur *userRepository) CheckUserPreference(userID uint64, preferenceName string) (bool, error) {
	var count int64
	err := ur.DB.Model(&domain.UserPreference{}).Where("user_id = ? AND preference_name = ?", userID, preferenceName).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ur *userRepository) CheckUserPreferenceByID(userID uint64, preferenceID uint64) (bool, error) {
	var count int64
	err := ur.DB.Model(&domain.UserPreference{}).Where("user_id = ? AND id = ?", userID, preferenceID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ur *userRepository) AddUserPreference(userID uint64, preferenceID int) error {
	preference := domain.UserPreference{
		UserID:       int(userID),
		PreferenceID: preferenceID,
	}
	return ur.DB.Create(&preference).Error
}

func (ur *userRepository) EditUserPreference(userID uint64, preferenceID uint64, newPreferenceName string) error {
	return ur.DB.Model(&domain.UserPreference{}).Where("user_id = ? AND id = ?", userID, preferenceID).Update("preference_name", newPreferenceName).Error
}

func (ur *userRepository) DeleteUserPreference(userID uint64, preferenceID uint64) error {
	return ur.DB.Delete(&domain.UserPreference{}, "user_id = ? AND id = ?", userID, preferenceID).Error
}

func (ur *userRepository) GetUserPreferences(userID uint64) ([]string, error) {
	var preferences []string
	err := ur.DB.Model(&domain.UserPreference{}).Where("user_id = ?", userID).Pluck("preference_id ", &preferences).Error
	if err != nil {
		return nil, err
	}
	return preferences, nil
}

func (ur *userRepository) GetUserInterests(userID uint64) ([]string, error) {
	var interests []string
	fmt.Println("id", userID)
	err := ur.DB.Model(&domain.UserInterest{}).Where("user_id = ?", userID).Pluck("interest_id", &interests).Error
	if err != nil {
		return nil, err
	}
	return interests, nil
}

//	func (ur *userRepository) FollowUser(senderID uint64, userID uint64) ([]string, error) {
//		var followReq domain.FollowUSer
//		fmt.Println("id", userID)
//		err := ur.DB.Model(&domain.UserInterest{}).Where("user_id = ?", userID).Pluck("interest_id", &interests).Error
//		if err != nil {
//			return nil, err
//		}
//		return interests, nil
//	}
func (ur *userRepository) CheckConnectionRequestExist(senderID, userID uint) (bool, error) {
	var connection domain.Connections
	result := ur.DB.Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", senderID, userID, userID, senderID).First(&connection)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// AddConnection adds a connection between senderID and userID
func (ur *userRepository) AddConnection(senderID, userID uint) error {
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
func (ur *userRepository) BlockConnection(senderID, userID uint) error {
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
func (ur *userRepository) CheckFriends(senderID, userID uint) (bool, error) {
	var connection domain.Connections
	result := ur.DB.Where("(user_id = ? AND friend_id = ?) AND status = 'friends'", senderID, userID).First(&connection)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// AddConnectionRequest adds a connection request between senderID and userID
func (ur *userRepository) AddConnectionRequest(senderID, userID uint) error {
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
func (ur *userRepository) SaveMessage(message *domain.UserMessage) (uint, error) {
	result := ur.DB.Create(message)
	if result.Error != nil {
		return 0, result.Error
	}
	return message.ID, nil
}

func (ur *userRepository) SaveMedia(media *domain.Media) error {
	fmt.Println("repohiii")
	fmt.Println(media)
	result := ur.DB.Create(media)

	if result.Error != nil {
		return result.Error
	}
	fmt.Println("nill")
	return nil
}
func (ur *userRepository) GetMessages(receiverID, senderID uint64) ([]models.UserMessage, error) {
	var messages []models.UserMessage

	// Preload media for each message
	if err := ur.DB.Preload("Media").Where("(recipent_id = ? AND sender_id = ?) OR ( recipent_id = ? AND sender_id = ?)", receiverID, senderID, senderID, receiverID).Find(&messages).Error; err != nil {
		return nil, errors.New("error retrieving messages")
	}

	return messages, nil
}
func (ur *userRepository) GetConnections(userID uint) ([]models.UserDetails, error) {
	// Query the database to fetch connections for the specified user ID
	var connections []models.UserDetails
	err := ur.DB.Table("connections").
		Joins("INNER JOIN users ON connections.friend_id = users.id").
		Select("users.id, users.firstname, users.lastname, users.email, users.phone").
		Where("connections.user_id = ? AND connections.status = 'friends'", userID).Or("connections.friend_id = ? AND connections.status = 'friends'", userID).
		Find(&connections).Error
	if err != nil {
		return nil, err
	}
	return connections, nil
}
