package tel

import (
	"context"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/telemetry"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var NRApp *newrelic.Application

var Initialize = func(cfg config.Configs) {
	NRApp = telemetry.Initialize(cfg)
}

var StartTransaction = telemetry.StartNRTransaction
var EndTransaction = telemetry.EndNRTransaction

var APMMiddleware = telemetry.Middleware

var TraceFromContext = func(ctx context.Context) (traceId string, spanId string) {
	data, _ := telemetry.TraceFromContext(ctx)
	return data.TraceId, data.SpanId
}

// SpanFuncWrapper Embrace your function with a New Relic segment/span
/* use:
func Foo() {
	defer SpanFuncWrapper("<TransactionName>")()
	...
}
*/
var SpanFuncWrapper = telemetry.NRSegmentFuncWrapper
