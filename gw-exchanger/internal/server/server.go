package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/chtozamm/javacode-final/gw-exchanger/internal/config"
	"github.com/chtozamm/javacode-final/gw-exchanger/internal/storage"
	logging "github.com/chtozamm/javacode-final/gw-exchanger/pkg/logs"
	pb "github.com/chtozamm/javacode-final/proto-exchange/exchange"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExchangeServiceServer
	log *logging.Logger
	cfg *config.Config
	db  storage.Repository
}

func NewServer(logger *logging.Logger, cfg *config.Config, db storage.Repository) *server {

	srv := &server{
		log: logger,
		cfg: cfg,
		db:  db,
	}

	return srv
}

func (s *server) Start() {
	// Create TCP listener
	lis, err := net.Listen("tcp", net.JoinHostPort(s.cfg.ServerHost, s.cfg.ServerPort))
	if err != nil {
		s.log.Fatal().Err(err).Msg("Failed to create TCP listener")
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterExchangeServiceServer(grpcServer, s)

	// Start gRPC server
	s.log.Info().Msgf("Starting gRPC server on %s:%s", s.cfg.ServerHost, s.cfg.ServerPort)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			s.log.Fatal().Err(err).Msg("Failed to start gRPC server")
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.log.Info().Msg("Shutting down gRPC server")
	grpcServer.GracefulStop()
}
