package interfaces

import "github.com/bibin-zoz/api-gateway/pkg/utils/models"

type AdminClient interface {
	AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error)
	GetIntrests() ([]models.Intrests, error)
	GetPreferences() ([]models.Preferences, error)
	AddInterest(interestName string) (models.InterestResponse, error)
	EditInterest(interestID int, newInterestName string) (models.InterestResponse, error)
	DeleteInterest(interestID int) error
	AddPreference(preferenceName string) (models.PreferenceResponse, error)
	EditPreference(preferenceID int, newPreferenceName string) (models.PreferenceResponse, error)
	DeletePreference(preferenceID int) error
}
