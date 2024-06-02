package usecase

import (
	"errors"

	"github.com/bibin-zoz/social-match-connections-svc/pkg/config"
	interfaces "github.com/bibin-zoz/social-match-connections-svc/pkg/repository/interface"
	services "github.com/bibin-zoz/social-match-connections-svc/pkg/usecase/interface"
	"github.com/bibin-zoz/social-match-connections-svc/pkg/utils/models"
)

type connectionUseCase struct {
	ConnectionRepository interfaces.ConnectionRepository
	Config               config.Config
}

func NewconnectionUseCase(repository interfaces.ConnectionRepository, config config.Config) services.ConnectionUseCase {
	return &connectionUseCase{
		ConnectionRepository: repository,
		Config:               config,
	}
}

func (uu *connectionUseCase) FollowUser(senderID, userID int64) error {
	reqExist, err := uu.ConnectionRepository.CheckConnectionRequestExist(uint(senderID), uint(userID))
	if err != nil {
		return errors.New("failed to check connection request")
	}

	if reqExist {
		// If a connection request exists, add them as friends
		err := uu.ConnectionRepository.AddConnection(uint(senderID), uint(userID))
		if err != nil {
			return errors.New("failed to add connection as friend")
		}
		return nil
	}

	// Check if they are already friends
	areFriends, err := uu.ConnectionRepository.CheckFriends(uint(senderID), uint(userID))
	if err != nil {
		return errors.New("failed to check if users are already friends")
	}

	if areFriends {
		return errors.New("users are already friends")
	}

	// If no connection request exists and they are not friends, send a new request
	err = uu.ConnectionRepository.AddConnectionRequest(uint(senderID), uint(userID))
	if err != nil {
		return errors.New("failed to send connection request")
	}

	return nil
}
func (uu *connectionUseCase) BlockConnection(senderID, userID int64) error {
	reqExist, err := uu.ConnectionRepository.CheckConnectionRequestExist(uint(senderID), uint(userID))
	if err != nil {
		return errors.New("no request found")
	}

	if reqExist {

		err := uu.ConnectionRepository.BlockConnection(uint(senderID), uint(userID))
		if err != nil {
			return errors.New("failed to block user")
		}
		return nil
	}

	return nil
}
func (ur *connectionUseCase) GetConnections(userID uint64) ([]models.UserDetails, error) {
	connections, err := ur.ConnectionRepository.GetConnections(uint(userID))
	if err != nil {
		return nil, errors.New("error retrieving connections")
	}
	return connections, nil
}

// func (uu *connectionUseCase) UnblockConnection(senderID, userID int64) error {
// 	reqExist, err := uu.ConnectionRepository.CheckConnectionRequestExist(uint(senderID), uint(userID))
// 	if err != nil {
// 		return errors.New("no request found")
// 	}

// 	if reqExist {

// 		err := uu.ConnectionRepository.UnblockConnection(uint(senderID), uint(userID))
// 		if err != nil {
// 			return errors.New("failed to block user")
// 		}
// 		return nil
// 	}

//		return nil
//	}
func (ur *connectionUseCase) CheckUserConnection(userID int64, recvID int64) (bool, error) {
	connections, err := ur.ConnectionRepository.CheckUserConnection((userID), (recvID))
	if err != nil {
		return false, errors.New("error retrieving connections")
	}
	return connections, nil
}
