package client

import pb "github.com/bibin-zoz/social-match-userauth-svc/pkg/pb/connections"

type ConnectionClientInterface interface {
	FollowUser(userID, senderID int64) (int64, error)
	BlockUserConnection(userID, senderID int64) (int64, error)
	GetConnections(userID uint64) ([]*pb.UserDetails, error)
}
