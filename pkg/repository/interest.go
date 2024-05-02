package repository

import (
	"errors"
	"strconv"

	"github.com/bibin-zoz/social-match-admin-svc/pkg/domain"
	interfaces "github.com/bibin-zoz/social-match-admin-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type interestRepository struct {
	DB *gorm.DB
}

func NewInterestRepository(DB *gorm.DB) interfaces.InterestRepository {
	return &interestRepository{
		DB: DB,
	}
}

func (ir *interestRepository) CheckInterest(interestID string) (bool, error) {
	var interest domain.Interest
	id, _ := strconv.Atoi(interestID)
	res := ir.DB.Where(&domain.Interest{ID: id}).First(&interest)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil // Interest does not exist
		}
		return false, res.Error // Other errors
	}
	return true, nil // Interest exists
}

// Other repository methods related to interests can be implemented here...
