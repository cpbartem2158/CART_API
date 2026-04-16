package db

import (
	"context"
	"embed"
	"fmt"

	"github.com/cpbartem2158/CART_API/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Connect(ctx context.Context, cfg config.DatabaseConfig) (*sqlx.DB, error) {

	configsDB := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d  sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port,
	)

	db, err := sqlx.ConnectContext(ctx, "postgres", configsDB)
	if err != nil {
		return nil, err
	}

	if err := runMigrations(db); err != nil {
		return nil, err
	}
	return db, nil
}

func runMigrations(db *sqlx.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		return err
	}
	return nil
}
