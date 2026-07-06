package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(cfg *Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.DB_URL)
	if err != nil {
		return nil, fmt.Errorf("Error to conect to database: %v", err)
	}

	return pool, nil
}
