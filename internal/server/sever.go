package server

import (
	"fmt"
	auth "github.com/alexandermatseev/go_auth/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const Port = 5001

type Server struct {
	grps     *grpc.Server
	listener net.Listener
}

func Init() (Server, error) {
	server := Server{}
	server.grps = grpc.NewServer()
	reflection.Register(server.grps)
	auth.RegisterUserAuthServer(server.grps, &Routes{})
	return server, nil
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	if err != nil {
		return err
	}
	s.listener = listener
	log.Printf("start server on %d port\n", Port)
	if err := s.grps.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() (Server, error) {
	s.grps.Stop()
	if err := s.listener.Close(); err != nil {
		return *s, err
	}
	return *s, nil
}
