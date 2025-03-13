package main

import (
	"flag"
	"log"
	"net"

	"github.com/chtozamm/javacode-final/gw-exchanger/internal/config"
	"github.com/chtozamm/javacode-final/gw-exchanger/internal/server"
	pb "github.com/chtozamm/javacode-final/proto-exchange/exchange"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	// Чтение аргументов командной строки
	var path string
	flag.StringVar(&path, "c", "", "Path to configuration file")
	flag.Parse()

	if path != "" {
		if err := godotenv.Load(path); err != nil {
			log.Fatalf("Error loading env from file: %v", err)
		}
	}
}

func main() {
	// Загрузка конфигурации
	cfg := config.NewConfig()

	// Создание TCP слушателя
	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.ServerHost, cfg.ServerPort))
	if err != nil {
		// TODO: handle error
	}

	// Создание gRPC сервера
	s := grpc.NewServer()
	pb.RegisterExchangeServiceServer(s, &server.Server{})

	// Запуск gRPC сервера
	log.Printf("Starting gRPC server on %s:%s", cfg.ServerHost, cfg.ServerPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
