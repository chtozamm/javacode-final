package main

import (
	"flag"
	"log"
	"os"

	"github.com/chtozamm/javacode-final/gw-exchanger/internal/config"
	"github.com/chtozamm/javacode-final/gw-exchanger/internal/server"
	"github.com/chtozamm/javacode-final/gw-exchanger/internal/storage/postgres"
	logging "github.com/chtozamm/javacode-final/gw-exchanger/pkg/logs"
	"github.com/joho/godotenv"
)

func init() {
	// Read command-line arguments
	var path string
	flag.StringVar(&path, "c", "", "Path to configuration file")
	flag.Parse()

	// Load environmental variables from file
	if path != "" {
		if err := godotenv.Load(path); err != nil {
			log.Fatalf("Error loading env from file: %v", err)
		}
	}
}

func main() {
	// Load configuration
	cfg := config.Load()

	// Set up logger
	logger := logging.New(os.Stdout, cfg.LogLevel)

	// Connect to database
	db, err := postgres.NewConnector(cfg.Storage, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// Start gRPC server
	server.NewServer(logger, cfg, db).Start()
}
