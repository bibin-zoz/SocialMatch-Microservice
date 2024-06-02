package repository

import (
	"errors"
	"fmt"

	"github.com/bibin-zoz/social-match-admin-svc/pkg/domain"
	interfaces "github.com/bibin-zoz/social-match-admin-svc/pkg/repository/interface"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/utils/models"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ad *adminRepository) AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error) {
	var model models.AdminDetailsResponse
	if err := ad.DB.Raw("INSERT INTO admins (firstname,lastname,email,password) VALUES (?, ?, ? ,?) RETURNING id,firstname,lastname,email", adminDetails.Firstname, adminDetails.Lastname, adminDetails.Email, adminDetails.Password).Scan(&model).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}
	return model, nil
}
func (ad *adminRepository) CheckAdminExistsByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	res := ad.DB.Where(&domain.Admin{Email: email}).First(&admin)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Admin{}, res.Error
	}
	return &admin, nil
}

func (ad *adminRepository) FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error) {
	var user models.AdminSignUp
	err := ad.DB.Raw("SELECT * FROM admins WHERE email=? ", admin.Email).Scan(&user).Error
	if err != nil {
		return models.AdminSignUp{}, errors.New("error checking user details")
	}
	return user, nil
}
func (ad *adminRepository) FetchInterests() ([]domain.Interest, error) {
	var Interest []domain.Interest
	if err := ad.DB.Find(&Interest).Error; err != nil {
		return nil, err
	}

	return Interest, nil
}
func (ad *adminRepository) FetchPreference() ([]models.Preference, error) {
	var Preference []models.Preference
	if err := ad.DB.Find(&Preference).Error; err != nil {
		return nil, err
	}
	fmt.Println("Preference")
	return Preference, nil
}
func (ad *adminRepository) AddInterest(interestName string) (int64, error) {
	// Implement the logic to add interest to the database
	// For example:
	interest := domain.Interest{InterestName: interestName}
	if err := ad.DB.Create(&interest).Error; err != nil {
		return 0, err
	}
	return int64(interest.ID), nil
}

func (ad *adminRepository) EditInterest(id int64, interestName string) error {
	// Implement the logic to edit interest in the database
	// For example:
	if err := ad.DB.Model(&domain.Interest{}).Where("id = ?", id).Update("interest_name", interestName).Error; err != nil {
		return err
	}
	return nil
}

func (ad *adminRepository) DeleteInterest(id int64) error {
	// Implement the logic to delete interest from the database
	// For example:
	if err := ad.DB.Delete(&domain.Interest{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (ad *adminRepository) AddPreference(preferenceName string) (int64, error) {
	// Implement the logic to add preference to the database
	// For example:
	preference := models.Preference{PreferenceName: preferenceName}
	if err := ad.DB.Create(&preference).Error; err != nil {
		return 0, err
	}
	return int64(preference.ID), nil
}

func (ad *adminRepository) EditPreference(id int64, preferenceName string) error {
	// Implement the logic to edit preference in the database
	// For example:
	if err := ad.DB.Model(&models.Preference{}).Where("id = ?", id).Update("preference_name", preferenceName).Error; err != nil {
		return err
	}
	return nil
}

func (ad *adminRepository) DeletePreference(id int64) error {
	// Implement the logic to delete preference from the database
	// For example:
	if err := ad.DB.Delete(&models.Preference{}, id).Error; err != nil {
		return err
	}
	return nil
}
