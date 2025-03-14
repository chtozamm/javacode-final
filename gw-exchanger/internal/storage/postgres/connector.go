package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/chtozamm/javacode-final/gw-exchanger/internal/cache"
	"github.com/chtozamm/javacode-final/gw-exchanger/internal/config"
	logging "github.com/chtozamm/javacode-final/gw-exchanger/pkg/logs"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Connector struct {
	Client *pgxpool.Pool
	log    *logging.Logger
	cfg    config.StorageConfig
	cache  *cache.Cache
}

func NewConnector(cfg config.StorageConfig, logger *logging.Logger) (*Connector, error) {
	cache := cache.New(time.Duration(cfg.CacheExpiration) * time.Second)
	c := &Connector{
		log:   logger,
		cfg:   cfg,
		cache: cache,
	}

	err := c.Start()
	if err != nil {
		return nil, err
	}

	go c.cache.Cleanup()

	return c, nil

}

func (c *Connector) Start() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create connection pool
	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		c.cfg.PostgresUsername,
		c.cfg.PostgresPassword,
		c.cfg.PostgresHost,
		c.cfg.PostgresPort,
		c.cfg.PostgresDatabase,
	))
	if err != nil {
		return fmt.Errorf("failed to create PostgreSQL connection pool: %v", err)
	}

	// Check the database connection
	err = pool.Ping(ctx)
	if err != nil {
		return fmt.Errorf("database unreachable: %v", err)
	}

	c.Client = pool
	c.log.Info().Msgf("Connected with PostgreSQL database %s:%s", c.cfg.PostgresHost, c.cfg.PostgresPort)
	return nil
}

func (c *Connector) Stop() error {
	c.Client.Close()
	return nil
}
