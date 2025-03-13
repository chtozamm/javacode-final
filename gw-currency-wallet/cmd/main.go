package main

import (
	"flag"
	"log"
	"net"

	"github.com/chtozamm/javacode-final/gw-currency-wallet/internal/config"
	"github.com/chtozamm/javacode-final/gw-currency-wallet/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	// Создание сервера
	r := gin.Default()

	// Регистрация обработчиков
	handler.Register(r)

	// Запуск сервера
	if err := r.Run(net.JoinHostPort(cfg.ServerHost, cfg.ServerPort)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
