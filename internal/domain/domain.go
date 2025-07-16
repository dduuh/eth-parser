package domain

type Addresses struct {
	Id         int    `json:"id" db:"id"`
	Address    string `json:"address" db:"address"`
	PrivateKey string `json:"privateKey" db:"private_key"`
}