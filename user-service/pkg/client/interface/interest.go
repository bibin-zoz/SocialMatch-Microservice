package client

type InterestClientInterface interface {
	CheckInterest(interestID string) (bool, error)
}
