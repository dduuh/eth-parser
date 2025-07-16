package domain

import "math/big"

type Addresses struct {
	Id         int    `json:"-" db:"id"`
	Address    string `json:"address" db:"address"`
	PrivateKey string `json:"privateKey" db:"private_key"`
}

type Transaction struct {
	From   string   `json:"from"`
	To     string   `json:"to"`
	Amount *big.Int `json:"amount"`
}
