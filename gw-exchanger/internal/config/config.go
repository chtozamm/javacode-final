package config

import (
	"github.com/chtozamm/javacode-final/gw-exchanger/pkg/env"
)

type Config struct {
	ServerHost string
	ServerPort string

	ClientTimeout      int
	ServerReadTimeout  int
	ServerWriteTimeout int

	LogLevel string

	Storage StorageConfig
}

type StorageConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string
	PostgresUsername string
	PostgresPassword string
	CacheExpiration  int
}

func Load() *Config {
	return newConfig()
}

func newConfig() *Config {
	storage := StorageConfig{
		PostgresHost:     env.GetEnv("DB_HOST", DefaultDBHost),
		PostgresPort:     env.GetEnv("DB_PORT", DefaultDBPort),
		PostgresDatabase: env.GetEnv("DB_NAME", DefaultDBName),
		PostgresUsername: env.GetEnv("DB_USERNAME", DefaultDBUsername),
		PostgresPassword: env.GetEnv("DB_PASSWORD", DefaultDBPassword),
		CacheExpiration:  env.GetEnvInt("CACHE_EXPIRATION", DefaultCacheExpiration),
	}

	config := &Config{
		ServerHost:         env.GetEnv("SERVER_HOST", DefaultHost),
		ServerPort:         env.GetEnv("SERVER_PORT", DefaultPort),
		ServerReadTimeout:  env.GetEnvInt("SERVER_READ_TIMEOUT", DefaultServerReadTimeout),
		ServerWriteTimeout: env.GetEnvInt("SERVER_WRITE_TIMEOUT", DefaultServerWriteTimeout),
		LogLevel:           env.GetEnv("LOG_LEVEL", DefaultLogLevel),
		Storage:            storage,
	}

	return config
}
