package interfaces

type InterestRepository interface {
	CheckInterest(interestID string) (bool, error)
	CheckInterestByName(interest string) (bool, error)
}
