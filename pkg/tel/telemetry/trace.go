package telemetry

import (
	"context"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/data"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func TraceFromContext(ctx context.Context) (data.TraceInfo, bool) {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		return data.TraceInfo{}, false
	}
	return data.TraceInfo{
		TraceId: txn.GetTraceMetadata().TraceID,
		SpanId:  txn.GetTraceMetadata().SpanID,
	}, true
}

func StartNRTransaction(ctx context.Context, name string) context.Context {
	tnx := App.StartTransaction(name)
	return newrelic.NewContext(ctx, tnx)
}

func EndNRTransaction(ctx context.Context, name string) {
	tnx := newrelic.FromContext(ctx)
	if tnx != nil && tnx.Name() == name {
		tnx.End()
	}
}

func NRSegmentFuncWrapper(ctx context.Context, name string) func() {
	txn := newrelic.FromContext(ctx)
	seg := txn.StartSegment(name)
	return func() {
		seg.End()
	}
}
