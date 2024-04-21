package interfaces

import (
	"github.com/bibin-zoz/social-match-admin-svc/pkg/domain"
	"github.com/bibin-zoz/social-match-admin-svc/pkg/utils/models"
)

type AdminUseCase interface {
	AdminSignUp(admindeatils models.AdminSignUp) (*domain.TokenAdmin, error)
	LoginHandler(adminDetails models.AdminLogin) (*domain.TokenAdmin, error)
}
