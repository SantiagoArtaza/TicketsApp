package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return dbpool, nil
}
