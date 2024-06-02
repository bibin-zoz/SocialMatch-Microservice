package service

import (
	"context"
	"fmt"

	pb "github.com/bibin-zoz/social-match-connections-svc/pkg/pb/connections"

	interfaces "github.com/bibin-zoz/social-match-connections-svc/pkg/usecase/interface"
)

type ConnectionServer struct {
	connectionUseCase interfaces.ConnectionUseCase
	pb.UnimplementedConnectionsServer
}

func NewUserServer(useCase interfaces.ConnectionUseCase) pb.ConnectionsServer {

	return &ConnectionServer{
		connectionUseCase: useCase,
	}

}
func (a *ConnectionServer) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.FollowUserResponce, error) {
	err := a.connectionUseCase.FollowUser(req.Senderid, req.Userid)
	if err != nil {
		return &pb.FollowUserResponce{}, err
	}
	return &pb.FollowUserResponce{
		Status: 201,
	}, nil
}

func (a *ConnectionServer) GetConnections(ctx context.Context, req *pb.GetConnectionsRequest) (*pb.GetConnectionsResponse, error) {
	// Call the use case method to get connections for the specified user ID
	connections, err := a.connectionUseCase.GetConnections(req.UserId)
	if err != nil {
		return nil, err
	}

	// Convert the returned connections to protobuf format and return
	var pbConnections []*pb.UserDetails
	for _, connection := range connections {
		pbConnection := &pb.UserDetails{
			Id: uint64(connection.ID),
		}
		pbConnections = append(pbConnections, pbConnection)
	}
	fmt.Println("userdetails", pbConnections)
	return &pb.GetConnectionsResponse{
		Status:      201,
		UserDetails: pbConnections,
	}, nil
}
func (a *ConnectionServer) BlockConnection(ctx context.Context, req *pb.BlockConnectionRequest) (*pb.BlockConnectionResponse, error) {
	err := a.connectionUseCase.BlockConnection(int64(req.SenderId), int64(req.UserId))
	if err != nil {
		return &pb.BlockConnectionResponse{}, err
	}
	return &pb.BlockConnectionResponse{
		Status: 201,
	}, nil
}
func (a *ConnectionServer) CheckUserConnection(ctx context.Context, req *pb.CheckUserConnectionReq) (*pb.ConnectionResponse, error) {
	friend, err := a.connectionUseCase.CheckUserConnection(int64(req.UserID), int64(req.ReceiverID))
	if err != nil {
		return &pb.ConnectionResponse{}, err
	}
	return &pb.ConnectionResponse{
		Connected: friend,
	}, nil
}
