package interfaces

import "context"

type InterestUseCase interface {
	CheckInterest(ctx context.Context, interestID string) (bool, error)
}
