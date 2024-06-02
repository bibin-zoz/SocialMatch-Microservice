package interfaces

type ConnectionClient interface {
	CheckUserConnection(userID int32, receiverID int32) (bool, error)
}
