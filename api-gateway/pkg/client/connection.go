package client

import (
	"context"
	"fmt"

	"github.com/bibin-zoz/api-gateway/pkg/client/interfaces"
	"github.com/bibin-zoz/api-gateway/pkg/config"
	Connection "github.com/bibin-zoz/api-gateway/pkg/pb/connections"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConnectionClient struct {
	// Client     pb.ConnectionServiceClient
	ConnectionClient Connection.ConnectionsClient
}

func NewConnectionServiceClient(cfg config.Config) interfaces.ConnectionClient {
	ConnectionGrpcConnection, err := grpc.Dial(cfg.ConnectionSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect to Connection service", err)
		return nil
	}
	ConnectionGrpcClient := Connection.NewConnectionsClient(ConnectionGrpcConnection)
	return &ConnectionClient{

		ConnectionClient: ConnectionGrpcClient,
	}

}
func (cc *ConnectionClient) CheckUserConnection(userID int32, receiverID int32) (bool, error) {
	req := &Connection.CheckUserConnectionReq{
		UserID:     userID,
		ReceiverID: receiverID,
	}

	resp, err := cc.ConnectionClient.CheckUserConnection(context.Background(), req)
	if err != nil {
		return false, fmt.Errorf("failed to check user connection: %v", err)
	}

	return resp.Connected, nil
}
