package scripts

import (
	"context"
	"fmt"
	"github.com/Gilmardealcantara/go-micro-svc/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"
)

func SetupPostgresContainer(ctx context.Context) (*postgres.PostgresContainer, error) {
	dbName := "go-service-template"
	dbUser := "go-service-template"
	dbPassword := "password"

	return postgres.Run(ctx,
		"postgres:16-alpine",
		//postgres.WithInitScripts(filepath.Join("testdata", "init-user-db.sh")),
		//postgres.WithConfigFile(filepath.Join("testdata", "my-postgres.conf")),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
}

func SetupDb(ctx context.Context, container *postgres.PostgresContainer) (*db.DB, error) {
	databaseUrl, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	dbCfg, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbCfg)
	if err != nil {
		return nil, err
	}

	err = db.MigrateUp(databaseUrl)
	if err != nil {
		return nil, err
	}

	return &db.DB{Pool: pool}, nil
}

func TearDownDb(connectedDb *db.DB, pgxContainer *postgres.PostgresContainer) error {
	println(connectedDb.Pool.Config().ConnString())
	err := db.MigrateDown(connectedDb.Pool.Config().ConnString())
	if err != nil {
		fmt.Printf("Failed to execute db migrate down: %s\n", err)
		return err
	}
	connectedDb.Pool.Close()

	if err := testcontainers.TerminateContainer(pgxContainer); err != nil {
		fmt.Printf("Failed to terminate postgres test container: %s\n", err)
		return err
	}
	return nil
}
