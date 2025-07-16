package rest

import (
	"crypto/ecdsa"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
)

type Dict map[string]any

func response(w http.ResponseWriter, statusCode int, msg any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp, err := json.Marshal(msg)
	if err != nil {
		logrus.Warnf("warn: cannot marshal the data: %v\n", err)
	}

	_, err = w.Write(resp)
	if err != nil {
		logrus.Warnf("warn: cannot write data to the conn: %v", err)
	}
}

func generateAddresses() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", ""
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", ""
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return address.String(), privateKey.X.String()
}
