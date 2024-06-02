package usecase

import (
	repository "github.com/bibin-zoz/social-match-admin-svc/pkg/repository/interface"
	services "github.com/bibin-zoz/social-match-admin-svc/pkg/usecase/interface"
)

type InterestUseCase struct {
	interestRepo repository.InterestRepository // Assuming you have an InterestRepository interface
}

func NewInterestUseCase(interestRep repository.InterestRepository) services.InterestUseCase {
	return &InterestUseCase{
		interestRepo: interestRep,
	}
}

func (uc *InterestUseCase) CheckInterest(interestID string) (bool, error) {
	exists, err := uc.interestRepo.CheckInterest(interestID)
	if err != nil {
		return false, err
	}
	return exists, nil
}
func (uc *InterestUseCase) CheckInterestbyName(interestName string) (bool, error) {
	exists, err := uc.interestRepo.CheckInterestByName(interestName)
	if err != nil {
		return false, err
	}
	return exists, nil
}
