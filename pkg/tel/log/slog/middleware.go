package slog

import (
	"context"
	"fmt"

	"log/slog"
	"net/http"
	"time"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	httputil "github.com/Gilmardealcantara/go-micro-svc/pkg/tel/http"
)

func Middleware(cfg config.Configs) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			start := time.Now()

			ctx := r.Context()
			rww := httputil.NewResponseWriterWrapper(rw)

			ctxBuilder := NewContextBuilder(ctx).AddRequestInfo(r)

			next.ServeHTTP(rww, r.WithContext(ctxBuilder.Build()))

			statusCode := rww.Code()
			if path == cfg.HealthCheckPath && statusCode == http.StatusOK {
				return
			}

			ctx = ctxBuilder.
				AddResponseInfo(rww).
				AddAccountInfo(rww).
				Build()

			latency := time.Since(start).String()
			message := fmt.Sprintf("%d | %12v | %s | %s  %s", statusCode, latency, r.RemoteAddr, r.Method, path)
			LogRequest(ctx, message, statusCode, cfg)
		})
	}
}

func LogRequest(ctx context.Context, message string, statusCode int, cfg config.Configs) {
	if statusCode >= 500 {
		slog.ErrorContext(ctx, message)
		return
	}
	if statusCode >= 400 {
		slog.WarnContext(ctx, message)
		return
	}
	if cfg.RequestInfoLog {
		slog.InfoContext(ctx, message)
		return
	}
	slog.DebugContext(ctx, message)
}
