package db

import (
	"context"
	"embed"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	driverName = "postgres"
)

type Settings struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func ConnectPostgres(ctx context.Context, settings Settings) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		settings.User,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.DBName,
	)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err = goose.SetDialect(driverName); err != nil {
		return nil, err
	}

	goose.SetBaseFS(embedMigrations)
	postgreDB := stdlib.OpenDBFromPool(pool)
	defer postgreDB.Close()

	if err = goose.Up(postgreDB, "migrations"); err != nil {
		return nil, err
	}

	return pool, nil
}
