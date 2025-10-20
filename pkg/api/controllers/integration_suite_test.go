package controllers_test

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"testing"

	"github.com/Gilmardealcantara/go-micro-svc/db"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/api"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/log"
	"github.com/Gilmardealcantara/go-micro-svc/scripts"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type IntegrationSuite struct {
	suite.Suite
	router       *http.ServeMux
	pgxContainer *postgres.PostgresContainer
	db           *db.DB
}

func (s *IntegrationSuite) SetupSuite() {
	var err error
	ctx := context.Background()
	s.pgxContainer, err = scripts.SetupPostgresContainer(ctx)
	if err != nil {
		s.FailNow("Failed to postgres test container test db", err)
	}

	s.db, err = scripts.SetupDb(ctx, s.pgxContainer)
	if err != nil {
		s.FailNow("Failed to setup test db", err)
	}

	cfg := config.NewConfig()

	l := slog.New(slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{}))
	slog.SetDefault(l)
	logger := log.New(l)

	s.router = api.SetupRouter(&cfg, s.db, logger)
}

func (s *IntegrationSuite) TearDownSuite() {
	_ = scripts.TearDownDb(s.db, s.pgxContainer)
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}
