package grpcserver

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
)

type IGRPCConfig interface {
	GetHost() string
	GetPort() string
	GetAddr() string
}

type GRPCConfig struct {
	host string
	port string
}

func NewConfig() IGRPCConfig {
	return &GRPCConfig{
		host: os.Getenv("GRPC_HOST"),
		port: os.Getenv("GRPC_PORT"),
	}
}

func (cfg *GRPCConfig) GetHost() string {
	return cfg.port
}

func (cfg *GRPCConfig) GetPort() string {
	return cfg.host
}

func (cfg *GRPCConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", cfg.host, cfg.port)
}

type GRPCServer struct {
	cfg      IGRPCConfig
	instance *grpc.Server
}

func NewGRPCServer(cfg IGRPCConfig) *GRPCServer {
	return &GRPCServer{
		cfg:      cfg,
		instance: grpc.NewServer(),
	}
}

func (s *GRPCServer) GetConfig() IGRPCConfig {
	return s.cfg
}

func (s *GRPCServer) GetInstance() *grpc.Server {
	return s.instance
}

func (s *GRPCServer) Run() error {
	l, err := net.Listen("tcp", s.cfg.GetAddr())
	if err != nil {
		return err
	}

	if err := s.instance.Serve(l); err != nil {
		return err
	}

	return nil
}
