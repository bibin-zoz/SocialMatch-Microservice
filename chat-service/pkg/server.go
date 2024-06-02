package server

import (
	"fmt"
	"net"

	"github.com/bibin-zoz/social-match-chat-svc/pkg/config"
	pb "github.com/bibin-zoz/social-match-chat-svc/pkg/pb/chat"
	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.ChatServiceServer) (*Server, error) {

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer()
	pb.RegisterChatServiceServer(newServer, server)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50057")
	return c.server.Serve(c.listener)
}
