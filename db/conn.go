// Package db/conn.go creates a connection to the database.
package db

import (
	"context"
	"log/slog"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/newrelic/go-agent/v3/integrations/nrpgx5"
)

type DB struct {
	Pool *pgxpool.Pool
}

func Conn(ctx context.Context, cfg config.Configs) (*DB, error) {
	dbCfg, err := pgxpool.ParseConfig(cfg.DatabaseUrl())
	if err != nil {
		return nil, err
	}

	dbCfg.BeforeConnect = func(_ context.Context, config *pgx.ConnConfig) error {
		config.Tracer = nrpgx5.NewTracer()
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbCfg)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	slog.InfoContext(ctx, "[Conn] Connected to database with success!!!")
	return &DB{Pool: pool}, nil
}

func ClosePool(connectedDB *DB) func() error {
	return func() error {
		if connectedDB.Pool != nil {
			connectedDB.Pool.Close()
		}
		return nil
	}
}
