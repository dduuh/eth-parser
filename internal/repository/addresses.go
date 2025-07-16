package repository

import (
	"context"
	"eth-parser/internal/domain"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Addresses struct {
	db *sqlx.DB
}

func NewAddresses(db *sqlx.DB) *Addresses {
	return &Addresses{
		db: db,
	}
}

func (a *Addresses) AddAddress(ctx context.Context, addr domain.Addresses) (domain.Addresses, error) {
	var createdAddr domain.Addresses
	
	query := `INSERT INTO addresses (id, address, private_key) VALUES ($1, $2, $3)`

	row := a.db.QueryRowContext(ctx, query, addr.Id, addr.Address, addr.PrivateKey)
	err := row.Scan(
		&createdAddr.Id,
		&createdAddr.Address,
		&createdAddr.PrivateKey)
	if err != nil {
		return domain.Addresses{}, fmt.Errorf("failed to add the address to the database: %w", err)
	}

	return createdAddr, nil
}

func (a *Addresses) GetAddresses(ctx context.Context) ([]domain.Addresses, error) {
	var addresses []domain.Addresses

	query := `SELECT id, address, private_key FROM addresses`
	err := a.db.SelectContext(ctx, &addresses, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all addresses: %w", err)
	}

	return addresses, nil
}