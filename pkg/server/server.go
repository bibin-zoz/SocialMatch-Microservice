package server

import (
	"fmt"
	"net"

	"github.com/bibin-zoz/social-match-admin-svc/pkg/config"
	pb "github.com/bibin-zoz/social-match-admin-svc/pkg/pb/admin"
	ipb "github.com/bibin-zoz/social-match-admin-svc/pkg/pb/interest"
	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.AdminServer, interestServer ipb.InterestServiceServer) (*Server, error) {

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer()
	pb.RegisterAdminServer(newServer, server)
	ipb.RegisterInterestServiceServer(newServer, interestServer)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50052")
	return c.server.Serve(c.listener)
}
