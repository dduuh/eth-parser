package main

import (
	"context"
	"eth-parser/internal/config"
	"eth-parser/internal/repository"
	"eth-parser/internal/repository/psql"
	"eth-parser/internal/service"
	"eth-parser/internal/transport/rest"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		logrus.Panicf("Configs error: %v\n", err)
	}

	postgres, err := psql.New(cfg)
	if err != nil {
		logrus.Panicf("PostgreSQL error: %v\n", err)
	}

	if err := postgres.Up(); err != nil {
		logrus.Panicf("Migrate error: %v\n", err)
	}

	addressRepo := repository.NewAddresses(postgres.DB())
	service := service.New(addressRepo)
	rest := rest.New(service)

	go func() {
		if err := rest.Run(context.Background(), cfg, rest.InitRoutes()); err != nil {
			logrus.Panicf("HTTP error: %v\n", err)
		}
	}()
}
