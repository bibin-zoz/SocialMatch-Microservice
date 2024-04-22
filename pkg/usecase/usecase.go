package usecase

import (
	"errors"
	"fmt"

	"github.com/bibin-zoz/social-match-admin-svc/pkg/domain"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/helper"
	interfaces "github.com/bibin-zoz/social-match-admin-svc/pkg/repository/interface"
	services "github.com/bibin-zoz/social-match-admin-svc/pkg/usecase/interface"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/utils/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUseCase(repository interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{

		adminRepository: repository,
	}
}
func (ad *adminUseCase) AdminSignUp(admin models.AdminSignUp) (*domain.TokenAdmin, error) {
	email, err := ad.adminRepository.CheckAdminExistsByEmail(admin.Email)
	fmt.Println(email)
	if err != nil {

		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email != nil {
		return &domain.TokenAdmin{}, errors.New("admin with this email is already exists")
	}
	hashPassword, err := helper.PasswordHash(admin.Password)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error in hashing password")
	}
	admin.Password = hashPassword
	admindata, err := ad.adminRepository.AdminSignUp(admin)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("could not add the user")
	}
	tokenString, err := helper.GenerateTokenAdmin(admindata)

	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	return &domain.TokenAdmin{
		Admin: admindata,
		Token: tokenString,
	}, nil
}

func (ad *adminUseCase) LoginHandler(admin models.AdminLogin) (*domain.TokenAdmin, error) {
	email, err := ad.adminRepository.CheckAdminExistsByEmail(admin.Email)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email == nil {
		return &domain.TokenAdmin{}, errors.New("email doesn't exist")
	}
	admindeatils, err := ad.adminRepository.FindAdminByEmail(admin)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}
	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(admindeatils.Password), []byte(admin.Password))
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("password not matching")
	}
	var adminDetailsResponse models.AdminDetailsResponse
	//  copy all details except password and sent it back to the front end
	err = copier.Copy(&adminDetailsResponse, &admindeatils)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	return &domain.TokenAdmin{
		Admin: adminDetailsResponse,
		Token: tokenString,
	}, nil
}
func (ad *adminUseCase) GetInterests() ([]domain.Interest, error) {
	interests, err := ad.adminRepository.FetchInterests()
	if err != nil {
		return []domain.Interest{}, errors.New("failed to fetch user details")
	}
	return interests, nil
}
func (ad *adminUseCase) GetPreferences() ([]models.Preference, error) {
	Preference, err := ad.adminRepository.FetchPreference()
	if err != nil {
		return []models.Preference{}, errors.New("failed to fetch user details")
	}
	return Preference, nil
}
func (ad *adminUseCase) AddInterest(interestName string) (int64, error) {
	// Call the repository method to add interest
	id, err := ad.adminRepository.AddInterest(interestName)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ad *adminUseCase) EditInterest(id int64, interestName string) error {
	// Call the repository method to edit interest
	err := ad.adminRepository.EditInterest(id, interestName)
	if err != nil {
		return err
	}
	return nil
}

func (ad *adminUseCase) DeleteInterest(id int64) error {
	// Call the repository method to delete interest
	err := ad.adminRepository.DeleteInterest(id)
	if err != nil {
		return err
	}
	return nil
}

func (ad *adminUseCase) AddPreference(preferenceName string) (int64, error) {
	// Call the repository method to add preference
	id, err := ad.adminRepository.AddPreference(preferenceName)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ad *adminUseCase) EditPreference(id int64, preferenceName string) error {
	// Call the repository method to edit preference
	err := ad.adminRepository.EditPreference(id, preferenceName)
	if err != nil {
		return err
	}
	return nil
}

func (ad *adminUseCase) DeletePreference(id int64) error {
	// Call the repository method to delete preference
	err := ad.adminRepository.DeletePreference(id)
	if err != nil {
		return err
	}
	return nil
}
