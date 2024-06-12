package interfaces

import (
	"github.com/bibin-zoz/social-match-admin-svc/pkg/domain"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)
	FetchPreference() ([]models.Preference, error)
	FetchInterests() ([]domain.Interest, error)

	// New methods for adding, editing, and deleting interests
	AddInterest(interestName string) (int64, error)
	EditInterest(id int64, interestName string) error
	DeleteInterest(id int64) error

	// New methods for adding, editing, and deleting preferences
	AddPreference(preferenceName string) (int64, error)
	EditPreference(id int64, preferenceName string) error
	DeletePreference(id int64) error
}
