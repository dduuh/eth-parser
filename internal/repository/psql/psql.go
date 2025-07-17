package psql

import (
	"eth-parser/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresDB struct {
	db *sqlx.DB
}

func New(cfg *config.Config) (*PostgresDB, error) {
	db, err := sqlx.Open("postgres", cfg.DSN())
	if err != nil {
		return &PostgresDB{}, fmt.Errorf("failed to open the postgres: %w", err)
	}

	if err := db.Ping(); err != nil {
		return &PostgresDB{}, fmt.Errorf("failed to ping the postgres: %w", err)
	}

	return &PostgresDB{
		db: db,
	}, nil
}

func (p *PostgresDB) Close() error {
	if err := p.db.Close(); err != nil {
		return fmt.Errorf("failed to close the db: %w", err)
	}

	return nil
}

func (p *PostgresDB) DB() *sqlx.DB {
	return p.db
}