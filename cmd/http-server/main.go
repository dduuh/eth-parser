package main

import (
	"context"
	"eth-parser/internal/config"
	"eth-parser/internal/repository"
	"eth-parser/internal/repository/psql"
	"eth-parser/internal/script"
	"eth-parser/internal/service"
	"eth-parser/internal/transport/rest"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
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

	tx := script.New()
	addressRepo := repository.NewAddresses(postgres.DB())
	service := service.New(addressRepo)
	rest := rest.New(service)

	go func() {
		if err := rest.Run(context.Background(), cfg, rest.InitRoutes()); err != nil {
			logrus.Panicf("HTTP error: %v\n", err)
		}
	}()

	addresses, err := addressRepo.GetAddresses(context.Background())
	if err != nil {
		logrus.Panicf("GetAddresses error: %v\n", err)
	}

	var parsedAddresses []common.Address
	for _, address := range addresses {
		parsedAddresses = append(parsedAddresses, common.HexToAddress(address.Address))
	}

	client, err := ethclient.Dial(cfg.ETH.NodeUrl)
	if err != nil {
		logrus.Panicf("ETH client error: %v\n", err)
	}

	go tx.MonitorBlocks(cfg, client, parsedAddresses)

	select {}
}
