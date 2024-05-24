package usecase

import (
	"errors"
	"fmt"
	"strings"
	"time"

	client "github.com/bibin-zoz/social-match-userauth-svc/pkg/client/interface"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/config"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/domain"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/helper"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/pb/chat"
	interfaces "github.com/bibin-zoz/social-match-userauth-svc/pkg/repository/interface"
	services "github.com/bibin-zoz/social-match-userauth-svc/pkg/usecase/interface"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepository   interfaces.UserRepository
	profilePhotoRepo interfaces.ProfilePhotoRepo
	Config           config.Config
	InterestClient   client.InterestClientInterface
	ConnectionClient client.ConnectionClientInterface
	ChatClient       client.ChatClientInterface
}

func NewUserUseCase(repository interfaces.UserRepository, profileRepo interfaces.ProfilePhotoRepo, config config.Config, interestClient client.InterestClientInterface, connectionClient client.ConnectionClientInterface, chatClient client.ChatClientInterface) services.UserUseCase {
	return &userUseCase{
		userRepository:   repository,
		profilePhotoRepo: profileRepo,
		Config:           config,
		InterestClient:   interestClient,
		ConnectionClient: connectionClient,
		ChatClient:       chatClient,
	}
}

func (ur *userUseCase) UsersSignUp(user models.UserSignup) (domain.TokenUser, error) {
	email, err := ur.userRepository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email)
	if err != nil {
		return domain.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return domain.TokenUser{}, errors.New("user with this email is already exists")
	}

	phone, err := ur.userRepository.CheckUserExistsByPhone(user.Number)
	fmt.Println(phone, nil)
	if err != nil {
		return domain.TokenUser{}, errors.New("error with server")
	}
	if phone != nil {
		return domain.TokenUser{}, errors.New("user with this phone is already exists")
	}
	if len(user.Password) < 6 {
		return domain.TokenUser{}, errors.New("password must be 6 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
	}

	// Validate email
	if !helper.ValidateEmail(user.Email) {
		return domain.TokenUser{}, errors.New("email validation failed.provide valid email")
	}
	if !strings.HasSuffix(strings.ToLower(user.Email), "@gmail.com") {
		return domain.TokenUser{}, errors.New("email must end with @gmail.com")
	}
	if !helper.ValidatePassword(user.Password) {
		return domain.TokenUser{}, errors.New("Password must be min length 6 and must contain alphabets and numbers")
	}
	hashPassword, err := helper.PasswordHash(user.Password)
	if err != nil {
		return domain.TokenUser{}, errors.New("error in hashing password")
	}
	user.Password = hashPassword
	userData, err := ur.userRepository.UserSignUp(user)
	if err != nil {
		return domain.TokenUser{}, errors.New("could not add the user")
	}
	accessToken, err := helper.GenerateAccessToken(userData)
	if err != nil {
		return domain.TokenUser{}, errors.New("couldn't create access token due to error")
	}
	refreshToken, err := helper.GenerateRefreshToken(userData)
	if err != nil {
		return domain.TokenUser{}, errors.New("couldn't create refresh token due to error")
	}
	return domain.TokenUser{
		User:         userData,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (ur *userUseCase) UsersLogin(user models.UserLogin) (domain.TokenUser, error) {
	email, err := ur.userRepository.CheckUserExistsByEmail(user.Email)
	if err != nil {
		return domain.TokenUser{}, errors.New("error with server")
	}
	if email == nil {
		return domain.TokenUser{}, errors.New("email doesn't exist")
	}
	userdeatils, err := ur.userRepository.FindUserByEmail(user)
	if err != nil {
		return domain.TokenUser{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userdeatils.Password), []byte(user.Password))
	if err != nil {
		return domain.TokenUser{}, errors.New("password not matching")
	}
	var userDetails models.UserDetails
	err = copier.Copy(&userDetails, &userdeatils)
	if err != nil {
		return domain.TokenUser{}, err
	}
	accessToken, err := helper.GenerateAccessToken(userDetails)
	if err != nil {
		return domain.TokenUser{}, errors.New("couldn't create accesstoken due to internal error")
	}
	refreshToken, err := helper.GenerateRefreshToken(userDetails)
	if err != nil {
		return domain.TokenUser{}, errors.New("counldn't create refreshtoken due to internal error")
	}
	return domain.TokenUser{
		User:         userDetails,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (ur *userUseCase) UserEditDetails(user models.UserSignup) (models.UserDetails, error) {
	email, err := ur.userRepository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email)
	if err != nil {
		return models.UserDetails{}, errors.New("error with server")
	}
	if email == nil {
		return models.UserDetails{}, errors.New("user not found")
	}

	if len(user.Password) < 6 {
		return models.UserDetails{}, errors.New("password must be 6 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
	}

	hashPassword, err := helper.PasswordHash(user.Password)
	if err != nil {
		return models.UserDetails{}, errors.New("error in hashing password")
	}
	user.Password = hashPassword
	userData, err := ur.userRepository.UserEditDetails(user)
	if err != nil {
		return models.UserDetails{}, errors.New("could not edit  user")
	}

	return userData, nil
}

func (ur *userUseCase) UserGenerateOtp(email string) (string, error) {
	mail, err := ur.userRepository.CheckUserExistsByEmail(email)
	fmt.Println(email)
	if err != nil {
		return "", errors.New("error with server")
	}
	if mail == nil {
		return "", errors.New("user not found")
	}
	otp, err := helper.SendOTP(email, ur.Config)
	if err != nil {
		return "", errors.New("failed to generate otp")
	}

	return otp, nil
}

func (ur *userUseCase) UserVerifyOtp(otp string, email string) (bool, error) {
	fmt.Println("otp usecase", otp)
	mail, err := ur.userRepository.CheckUserExistsByEmail(email)
	fmt.Println(email)
	if err != nil {
		return false, errors.New("error with server")
	}
	if mail == nil {
		return false, errors.New("user not found")
	}
	if !helper.VerifyOTP(otp, email) {
		return false, errors.New("invalid otp")
	}

	return true, nil
}

func (ur *userUseCase) GetAllUsers() ([]domain.User, error) {
	users, err := ur.userRepository.FetchAllUsers()
	if err != nil {
		return []domain.User{}, errors.New("failed to fetch user details")
	}
	return users, nil
}

func (ur *userUseCase) UpdateUserStatus(id int) (models.UserDetail, error) {
	user, err := ur.userRepository.GetUserByID(id)
	if err != nil {
		return models.UserDetail{}, errors.New("failed to fetch user status")
	}
	if user.Blocked {
		err := ur.userRepository.UnblockUser(uint(id))
		if err != nil {
			return models.UserDetail{}, errors.New("failed to block user")
		}
		user.Blocked = true

	} else {
		err := ur.userRepository.BlockUser(uint(id))
		if err != nil {
			return models.UserDetail{}, errors.New("failed to block user")
		}
		user.Blocked = false
	}
	return models.UserDetail{
		Firstname: user.Firstname,
		Email:     user.Email,
	}, nil
}

func (ur *userUseCase) AddUserInterest(userID uint64, interestID int) error {
	// Check if the user exists
	_, err := ur.userRepository.GetUserByID(int(userID))
	if err != nil {
		return errors.New("failed to get user details")
	}
	exist, err := ur.InterestClient.CheckInterest(string(interestID))
	if exist {
		return errors.New("invalid interest id")
	}

	// Check if the interest already exists for the user
	interestExists, err := ur.userRepository.CheckUserInterestByID(userID, uint64(interestID))
	if err != nil {
		return errors.New("failed to check user interest")
	}
	if interestExists {
		return errors.New("interest already exists for the user")
	}

	// Add the interest for the user
	err = ur.userRepository.AddUserInterest(userID, interestID)
	if err != nil {
		return errors.New("failed to add user interest")
	}

	return nil
}

func (ur *userUseCase) EditUserInterest(userID uint64, interestID uint64, newInterestName string) error {
	// Check if the user exists
	_, err := ur.userRepository.GetUserByID(int(userID))
	if err != nil {
		return errors.New("failed to get user details")
	}

	// Check if the interest exists for the user
	interestExists, err := ur.userRepository.CheckUserInterestByID(userID, interestID)
	if err != nil {
		return errors.New("failed to check user interest")
	}
	if !interestExists {
		return errors.New("interest does not exist for the user")
	}

	// Edit the user interest
	err = ur.userRepository.EditUserInterest(userID, interestID, newInterestName)
	if err != nil {
		return errors.New("failed to edit user interest")
	}

	return nil
}

func (ur *userUseCase) DeleteUserInterest(userID uint64, interestID uint64) error {
	// Check if the user exists
	_, err := ur.userRepository.GetUserByID(int(userID))
	if err != nil {
		return errors.New("failed to get user details")
	}
	// if user == nil {
	// 	return errors.New("user not found")
	// }

	// Check if the interest exists for the user
	interestExists, err := ur.userRepository.CheckUserInterestByID(userID, interestID)
	if err != nil {
		return errors.New("failed to check user interest")
	}
	if !interestExists {
		return errors.New("interest does not exist for the user")
	}

	// Delete the user interest
	err = ur.userRepository.DeleteUserInterest(userID, interestID)
	if err != nil {
		return errors.New("failed to delete user interest")
	}

	return nil
}
func (uu *userUseCase) AddUserPreference(userID uint64, preferenceID int) error {
	// Check if the user exists
	_, err := uu.userRepository.GetUserByID(int(userID))
	if err != nil {
		return errors.New("failed to get user details")
	}

	// Check if the preference already exists for the user
	preferenceExists, err := uu.userRepository.CheckUserPreferenceByID(userID, uint64(preferenceID))
	if err != nil {
		return errors.New("failed to check user preference")
	}
	if preferenceExists {
		return errors.New("preference already exists for the user")
	}

	// Add the preference for the user
	err = uu.userRepository.AddUserPreference(userID, preferenceID)
	if err != nil {
		return errors.New("failed to add user preference")
	}

	return nil
}

func (uu *userUseCase) EditUserPreference(userID uint64, preferenceID uint64, newPreferenceName string) error {
	// Check if the user exists
	_, err := uu.userRepository.GetUserByID(int(userID))
	if err != nil {
		return errors.New("failed to get user details")
	}

	// Check if the preference exists for the user
	preferenceExists, err := uu.userRepository.CheckUserPreferenceByID(userID, preferenceID)
	if err != nil {
		return errors.New("failed to check user preference")
	}
	if !preferenceExists {
		return errors.New("preference does not exist for the user")
	}

	// Edit the user preference
	err = uu.userRepository.EditUserPreference(userID, preferenceID, newPreferenceName)
	if err != nil {
		return errors.New("failed to edit user preference")
	}

	return nil
}

func (uu *userUseCase) DeleteUserPreference(userID uint64, preferenceID uint64) error {
	// Check if the user exists
	_, err := uu.userRepository.GetUserByID(int(userID))
	if err != nil {
		return errors.New("failed to get user details")
	}

	// Check if the preference exists for the user
	preferenceExists, err := uu.userRepository.CheckUserPreferenceByID(userID, preferenceID)
	if err != nil {
		return errors.New("failed to check user preference")
	}
	if !preferenceExists {
		return errors.New("preference does not exist for the user")
	}

	// Delete the user preference
	err = uu.userRepository.DeleteUserPreference(userID, preferenceID)
	if err != nil {
		return errors.New("failed to delete user preference")
	}

	return nil
}

func (uu *userUseCase) GetUserInterests(userID uint64) ([]string, error) {
	// Fetch the interests for the user
	interests, err := uu.userRepository.GetUserInterests(userID)
	if err != nil {
		return nil, errors.New("failed to fetch user interests")
	}
	return interests, nil
}

func (uu *userUseCase) GetUserPreferences(userID uint64) ([]string, error) {
	// Fetch the preferences for the user
	preferences, err := uu.userRepository.GetUserPreferences(userID)
	if err != nil {
		return nil, errors.New("failed to fetch user preferences")
	}
	return preferences, nil
}
func (uu *userUseCase) FollowUser(senderID, userID int64) error {
	_, err := uu.ConnectionClient.FollowUser(int64(senderID), int64(userID))
	if err != nil {
		return err
	}

	return nil
}
func (uu *userUseCase) BlockConnection(senderID, userID int64) error {
	_, err := uu.ConnectionClient.BlockUserConnection(int64(senderID), int64(userID))
	if err != nil {
		return err
	}

	return nil
}
func (ur *userUseCase) SendMessage(message *models.UserMessage) error {
	fmt.Println("usecase")
	// Assuming you have a repository method to save the message
	usermsg := &domain.UserMessage{
		SenderID:   message.SenderID,
		RecipentID: message.RecipentID,
		Content:    message.Content,
		CreatedAt:  message.CreatedAt,
	}
	var media []models.Media
	if message.Media != nil {

		for _, m := range message.Media {
			media = append(media, m)
			// media := &domain.Media{
			// 	Message_id: int(msgID),
			// 	Filename:   m.Filename,
			// }
			// fmt.Println("hi2")
			// err := ur.userRepository.SaveMedia(media)
			// if err != nil {
			// 	return errors.New("failed to save message")
			// }
		}

	}
	_, err := ur.ChatClient.SendMessage(uint32(usermsg.SenderID), uint32(usermsg.RecipentID), usermsg.Content, media)
	if err != nil {
		return err
	}
	fmt.Println("messaged saved successfully")
	return nil
}

// func (ur *userUseCase) ConsumeAndProcessMessages() {
// 	config := kafka.ReaderConfig{
// 		Brokers: []string{"localhost:9092"},
// 		GroupID: "1",
// 		Topic:   "chat",
// 	}

// 	consumer := kafka.NewReader(config)

// 	signals := make(chan os.Signal, 1)
// 	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
// 	defer signal.Stop(signals)

// 	// Handle signals
// 	go func() {
// 		sig := <-signals
// 		fmt.Printf("Received signal: %v\n", sig)
// 		fmt.Println("Shutting down...")

// 		// Close Kafka consumer
// 		if err := consumer.Close(); err != nil {
// 			fmt.Printf("Error closing Kafka consumer: %v\n", err)
// 		}

// 		os.Exit(0)
// 	}()

// 	// Continuously consume messages
// 	for {
// 		msg, err := consumer.ReadMessage(context.Background())
// 		if err != nil {
// 			fmt.Printf("Error reading message: %v\n", err)
// 			continue
// 		}

// 		var userMessage models.UserMessage
// 		if err := json.Unmarshal(msg.Value, &userMessage); err != nil {
// 			fmt.Printf("Failed to deserialize message: %v\n", err)
// 			continue
// 		}

// 		usermsg := &models.UserMessage{
// 			SenderID:   userMessage.SenderID,
// 			RecipentID: userMessage.RecipentID,
// 			Content:    userMessage.Content,
// 			CreatedAt:  userMessage.CreatedAt,
// 		}

// 		if err := ur.SendMessage(usermsg); err != nil {
// 			fmt.Printf("Failed to process message: %v\n", err)
// 		}
// 		fmt.Println("usermsg", userMessage)
// 	}
// }

func (ur *userUseCase) GetMessages(receiverID, senderID uint64) ([]models.UserMessage, error) {
	messages, err := ur.ChatClient.ReadMessages(uint32(receiverID), uint32(senderID))
	if err != nil {
		return nil, errors.New("error retrieving messages")
	}

	// Convert []*pb.Message to []models.UserMessage
	var userMessages []models.UserMessage
	for _, msg := range messages {
		userMessage := models.UserMessage{
			ID:         uint(msg.MessageId),
			RecipentID: uint(msg.ReceiverIdId),
			SenderID:   uint(msg.SenderId),
			Content:    msg.Content,
			CreatedAt:  time.Now(), // Assuming models.UserMessage.Timestamp is of the same type
			Media:      convertMedia(msg.Media),
		}
		userMessages = append(userMessages, userMessage)
	}
	return userMessages, nil
}

// Helper function to convert []*pb.Media to []models.Media
func convertMedia(pbMedia []*chat.Media) []models.Media {
	var media []models.Media
	for _, m := range pbMedia {
		media = append(media, models.Media{
			Filename: m.Filename,
			// Add other fields as necessary
		})
	}
	return media
}
func (ur *userUseCase) GetConnections(userID uint64) ([]models.UserDetails, error) {
	// Fetch connections from the ConnectionClient
	connections, err := ur.ConnectionClient.GetConnections(userID)
	if err != nil {
		return nil, errors.New("error retrieving connections")
	}

	// Convert []*connections.UserDetails to []models.UserDetails
	var userDetails []models.UserDetails
	for _, conn := range connections {
		userDetails = append(userDetails, models.UserDetails{
			ID: uint(conn.Id),
			// Name:      conn.Name,
			// Email:     conn.Email,
			// Number:    conn.Number,
			// Status:    conn.Status,
			// CreatedAt: conn.CreatedAt,
		})
	}

	return userDetails, nil
}
