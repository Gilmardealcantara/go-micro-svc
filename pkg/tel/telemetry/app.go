package telemetry

import (
	"log/slog"
	"os"
	"time"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var App *newrelic.Application

func Initialize(cfg config.Configs) *newrelic.Application {
	if cfg.NRDisabled == "true" {
		return nil
	}

	appName := cfg.AppName + "-" + cfg.AppEnv
	licenseKey := cfg.NRLicenceKey

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigEnabled(true),
		newrelic.ConfigAppLogEnabled(true),
		newrelic.ConfigAppLogMetricsEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigAppLogDecoratingEnabled(true),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigCodeLevelMetricsEnabled(true),
		newrelic.ConfigInfoLogger(os.Stdout),
	)

	if err != nil {
		slog.Error("[NewRelic] Error with NewRelic initialization", slog.Any("error", err))
		return app
	}

	err = app.WaitForConnection(time.Second * 5)
	if err != nil {
		slog.Error("[NewRelic] Error with NewRelic connection", slog.Any("error", err))
		return app
	}

	slog.Info("[NewRelic] NewRelic initialized with success!!")
	App = app
	return app
}
