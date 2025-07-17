package script

import (
	"bytes"
	"context"
	"encoding/json"
	"eth-parser/internal/config"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type Transaction struct {
	client *ethclient.Client
}

func New() *Transaction {
	return &Transaction{
		client: &ethclient.Client{},
	}
}

func (t *Transaction) MonitorBlocks(cfg *config.Config, client *ethclient.Client, addresses []common.Address) {
	headers := make(chan *types.Header)
	subscribe, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		logrus.Errorf("SubscribeNewHead() error: %v\n", err)
		return
	}

	for {
		select {
		case err := <-subscribe.Err():
			logrus.Errorf("<-subscribe.Err() error: %v\n", err)
			return
		case _ = <-headers:
			block, err := client.BlockByNumber(context.Background(), nil)
			if err != nil {
				continue
			}

			for _, tx := range block.Transactions() {
				for _, address := range addresses {
					if tx.To() != nil && *tx.To() == address {
						logrus.Infof("Отследили адрес транзакций: %s\n", address.Hex())
						telegramMessage := fmt.Sprintf("Отследили адрес транзакций: %s\n", address.Hex())

						err := sendTelegramMessage(
							cfg.Telegram.Token,
							cfg.Telegram.ChatId,
							telegramMessage)
						if err != nil {
							logrus.Panicf("Telegram error: %v\n", err)
							continue
						}
					}
				}
			}
		}
	}
}

func sendTelegramMessage(token string, chatId int64, msg string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	contentType := "application/json"

	body, err := json.Marshal(map[string]any{
		"chat_id": chatId,
		"text":    msg,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal the data: %w", err)
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to POST the request: %w", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			return
		}
	}()

	return nil
}

func extractMessageFromTx(tx *types.Transaction) (common.Address, error) {
	var signer types.Signer
	
	switch tx.Type() {
	case types.LegacyTxType:
		signer = types.NewEIP155Signer(tx.ChainId())
	}

	sender, err := signer.Sender(tx)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get the sender: %w", err)
	}

	return sender, nil
}
