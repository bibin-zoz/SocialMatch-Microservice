package interfaces

import (
	"github.com/bibin-zoz/social-match-admin-svc/pkg/domain"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)
	
}
