package psql

import (
	"errors"
	"eth-parser/internal/config"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (p *PostgresDB) Up() error {
	driver, err := postgres.WithInstance(p.db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to get driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://D:/Oracle/LETSGOOOOOO/eth-parser/migrations", "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to get migrate instance: %w", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logrus.Info("no migrations to apply")
		} else {
			return fmt.Errorf("failed to Up() migrations: %w", err)
		}
	}

	return nil
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