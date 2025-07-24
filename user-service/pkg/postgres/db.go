package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	URL             string        `env:"POSTGRES_URL,notEmpty"`
	MaxConns        int32         `env:"POSTGRES_MAX_CONNS" envDefault:"10"`
	MaxConnIdleTime time.Duration `env:"POSTGRES_MAX_CONN_IDLE" envDefault:"30s"`
	MaxConnLifetime time.Duration `env:"POSTGRES_MAX_CONN_LIFETIME" envDefault:"1h"`
}

func NewPostgres(cfg Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.MaxConns
	config.MaxConnIdleTime = cfg.MaxConnIdleTime
	config.MaxConnLifetime = cfg.MaxConnLifetime

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
