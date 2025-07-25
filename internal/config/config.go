package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		PostgresSQL PostgreSQLCfg
		HTTP        HTTPCfg
		ETH      ETHCfg
		Telegram    TelegramCfg
	}

	HTTPCfg struct {
		Port string `envconfig:"HTTP_PORT" default:"8081"`
	}

	PostgreSQLCfg struct {
		Host     string `envconfig:"POSTGRES_HOST"     default:"localhost"`
		Port     string `envconfig:"POSTGRES_PORT"     default:"5432"`
		User     string `envconfig:"POSTGRES_USER"     default:"postgres"`
		Password string `envconfig:"POSTGRES_PASSWORD" default:"Eth-Parser"`
		DBName   string `envconfig:"POSTGRES_DBNAME"   default:"postgres"`
		SSLMode  string `envconfig:"POSTGRES_SSLMODE"  default:"disable"`
	}

	ETHCfg struct {
		NodeUrl string `envconfig:"INFURA_PROJECT_ID" default:"wss://mainnet.infura.io/ws/v3/666904d95aad496a853ad440c0823bf2"`
	}

	TelegramCfg struct {
		Token  string `envconfig:"BOT_TOKEN" default:"7871111228:AAHquySw6jR9o2aDeMIWcyiFtK5BJfhOaVY"`
		ChatId int64  `envconfig:"CHAT_ID" default:"7512400345"`
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
		c.PostgresSQL.User,
		c.PostgresSQL.Password,
		c.PostgresSQL.Host,
		c.PostgresSQL.Port,
		c.PostgresSQL.DBName,
		c.PostgresSQL.SSLMode)
}
