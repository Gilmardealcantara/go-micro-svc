package tel

import (
	"log/slog"
	"net/http"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	islog "github.com/Gilmardealcantara/go-micro-svc/pkg/tel/log/slog"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/telemetry"
)

var InitializeWithSlog = func(cfg config.Configs) (*slog.Logger, error) {
	NRApp = telemetry.Initialize(cfg)
	return islog.Setup(cfg)
}

func SetupSlog(cfg config.Configs) (*slog.Logger, error) {
	return islog.Setup(cfg)
}

var Middleware = func(cfgs ...config.Func) func(next http.Handler) http.Handler {
	cfg := config.Builder(cfgs...)
	return islog.Middleware(cfg)
}
