package slog

import (
	"log/slog"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/data"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/telemetry"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrslog"
)

func Setup(cfg config.Configs) (*slog.Logger, error) {
	level := new(slog.LevelVar)
	level.Set(getLogModeFromConfig(cfg.LogLevel))

	options := slog.HandlerOptions{Level: level}
	fixedAttributes := []slog.Attr{
		slog.String("app_name", cfg.AppName),
		slog.String("app_version", cfg.DevopsCommit),
		slog.String("env", cfg.AppEnv),
		slog.String("namespace_name", cfg.NamespaceName),
		slog.String("pod_name", cfg.PodName),
	}

	var appHandler *appLoggerHandler

	if telemetry.App != nil {
		nrHandler := nrslog.JSONHandler(telemetry.App, cfg.LogOutput, &options)
		appHandler = NewHandler(nrHandler, fixedAttributes...)
	} else {
		jsonHandler := slog.NewJSONHandler(cfg.LogOutput, &options)
		appHandler = NewHandler(jsonHandler, fixedAttributes...)
	}
	logger := slog.New(appHandler)
	slog.SetDefault(logger)
	startupStatusInfo()
	return logger, nil
}

func getLogModeFromConfig(logLevel data.TLogLevel) slog.Level {
	d := map[string]slog.Level{
		data.LevelDebug: slog.LevelDebug,
		data.LevelError: slog.LevelError,
		data.LevelWarn:  slog.LevelWarn,
		data.LevelInfo:  slog.LevelInfo,
	}
	v, ok := d[logLevel]
	if !ok {
		return slog.LevelInfo
	}
	return v
}

func startupStatusInfo() {
	if telemetry.App != nil {
		slog.Info("[Setup] Structured Slog initialized WITH newrelic agent integration")
	} else {
		slog.Warn("[Setup] Structured Slog initialized WITHOUT newrelic agent integration")
	}
}
