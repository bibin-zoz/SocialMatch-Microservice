package interfaces

type InterestUseCase interface {
	CheckInterest(interestID string) (bool, error)
}
