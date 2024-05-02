package interfaces

type InterestRepository interface {
	CheckInterest(interestID string) (bool, error)
}
