package telemetry

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if App != nil {
			txn := App.StartTransaction(r.Pattern)
			defer txn.End()
			w = txn.SetWebResponse(w)
			txn.SetWebRequestHTTP(r)
			r = newrelic.RequestWithTransactionContext(r, txn)
		}
		next.ServeHTTP(w, r)
	})
}
