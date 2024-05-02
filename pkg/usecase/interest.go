package usecase

import (
	repository "github.com/bibin-zoz/social-match-admin-svc/pkg/repository/interface" // Import repository package
)

type InterestUseCase struct {
	interestRepo repository.InterestRepository // Assuming you have an InterestRepository interface
}

func NewInterestUseCase(interestRepo repository.InterestRepository) *InterestUseCase {
	return &InterestUseCase{
		interestRepo: interestRepo,
	}
}

func (uc *InterestUseCase) CheckInterest(interestID string) (bool, error) {
	exists, err := uc.interestRepo.CheckInterest(interestID)
	if err != nil {
		return false, err
	}
	return exists, nil
}
