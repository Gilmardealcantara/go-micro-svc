package log_test

import (
	"bytes"
	"context"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/log"

	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomLoggerConfiguration(t *testing.T) {
	buff := bytes.NewBuffer(nil)
	slogInstance := slog.New(slog.NewJSONHandler(buff, &slog.HandlerOptions{}))
	logger := log.New(slogInstance)

	logger.Info(context.TODO(), "my test logging statement")

	output := buff.String()
	assert.Contains(t, output, "\"level\":\"INFO\"")
	assert.Contains(t, output, "\"time\":\"")
	assert.Contains(t, output, "\"msg\":\"my test logging statement\"")
}
