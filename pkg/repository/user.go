package repository

import (
	"errors"
	"fmt"

	"github.com/bibin-zoz/social-match-userauth-svc/pkg/domain"
	interfaces "github.com/bibin-zoz/social-match-userauth-svc/pkg/repository/interface"
	"github.com/bibin-zoz/social-match-userauth-svc/pkg/utils/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}

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
