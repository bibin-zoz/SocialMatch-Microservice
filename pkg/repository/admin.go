package repository

import (
	"errors"

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

	return Preference, nil
}
