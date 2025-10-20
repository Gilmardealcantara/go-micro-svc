package slog

import (
	"context"
	"log/slog"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/telemetry"
)

type appLoggerHandler struct {
	slog.Handler
	fixedAttributes []slog.Attr
}

func NewHandler(handler slog.Handler, attrs ...slog.Attr) *appLoggerHandler {
	return &appLoggerHandler{handler, attrs}
}

func (h appLoggerHandler) Handle(ctx context.Context, record slog.Record) error {
	if len(h.fixedAttributes) > 0 {
		record.AddAttrs(h.fixedAttributes...)
	}

	if info, ok := telemetry.TraceFromContext(ctx); ok {
		record.Add(slog.String("trace.id", info.TraceId))
		record.Add(slog.String("span.id", info.SpanId))
	}

	if info, ok := HttpRequestInfoFromContext(ctx); ok {
		record.Add(slog.String("scope", "http_request"))
		record.Add(slog.String("request.client_service", info.ClientService))
		record.Add(slog.String("request.ip_address", info.IPAddress))
		record.Add(slog.String("request.method", info.Method))
		record.Add(slog.String("request.path", info.Path))
		record.Add(slog.String("request.user_agent", info.UserAgent))
		addString(&record, "request.origin", info.Origin)
		addString(&record, "request.xforwarded_host", info.XforwardedHost)
		addString(&record, "request.raw_query", info.RawQuery)
		addString(&record, "request.school_id", info.SchoolId)
		addString(&record, "request.account_id", info.AccountId)
		addString(&record, "request.user_id", info.UserId)
	}

	if info, ok := HttpResponseInfoFromContext(ctx); ok {
		record.Add(slog.Int("response.status_code", info.StatusCode))
		addString(&record, "response.body", info.Body)
		addString(&record, "response.error_msg", info.ErrorMsg)
	}

	if info, ok := AccountFromContext(ctx); ok && info != nil {
		record.Add("account", *info)
	}

	if logData, ok := CustomInfoFromContext(ctx); ok {
		for k, v := range logData {
			if v != nil && v != "" {
				record.Add(slog.Any(k, v))
			}
		}
	}

	return h.Handler.Handle(ctx, record)
}

//	func (a appLoggerHandler) Enabled(ctx context.ContextValues, level slog.Level) bool {
//		//TODO implement me
//		panic("implement me")
//	}

//func (a appLoggerHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
//	//TODO implement me
//	panic("implement me")
//}

//
//func (a appLoggerHandler) WithGroup(name string) slog.Handler {
//	//TODO implement me
//	panic("implement me")
//}

func addString(record *slog.Record, prop string, value string) {
	if value != "" {
		record.Add(slog.String(prop, value))
	}
}
