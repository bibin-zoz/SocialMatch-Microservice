package interfaces

import "github.com/bibin-zoz/api-gateway/pkg/utils/models"

type AdminClient interface {
	AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error)
	GetPreferences() ([]models.Intrests, error)
	GetIntrests() ([]models.Intrests, error)
}
