package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Postgres PostgreSQL
		HTTP HTTPCfg
	}

	HTTPCfg struct {
		Port string `envconfig:"HTTP_PORT" default:"8081"`
	}

	PostgreSQL struct {
		Host     string `envconfig:"POSTGRES_HOST"     default:"localhost"`
		Port     string `envconfig:"POSTGRES_PORT"     default:"5432"`
		User     string `envconfig:"POSTGRES_USER"     default:"postgres"`
		Password string `envconfig:"POSTGRES_PASSWORD" default:"Eth-Parser"`
		DBName   string `envconfig:"POSTGRES_DBNAME"   default:"postgres"`
		SSLMode  string `envconfig:"POSTGRES_SSLMODE"  default:"disable"`
	}
)

func Init() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return &Config{}, fmt.Errorf("failed to process configs: %w", err)
	}

	return &cfg, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.DBName,
		c.Postgres.SSLMode)
}