package tel

import (
	"context"

	"net/http"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/data"
	thttp "github.com/Gilmardealcantara/go-micro-svc/pkg/tel/http"
	islog "github.com/Gilmardealcantara/go-micro-svc/pkg/tel/log/slog"
)

type HttpClient = thttp.Client

var NewHttpClient = func(client *http.Client, cfg config.Configs) thttp.Client {
	return thttp.NewHttpClient(client, cfg)
}

type ResponseWriterWrapper = thttp.ResponseWriterWrapper

var NewResponseWriterWrapper = thttp.NewResponseWriterWrapper

// WriteHttpError adds error in request log record
var WriteHttpError = thttp.WriteHttpError

func WriteAccount(ctx context.Context, w http.ResponseWriter, acc data.Account) context.Context {
	if rww, ok := w.(thttp.ResponseWriterWrapper); ok {
		rww.WriteAccount(acc)
		return islog.NewContextBuilder(ctx).AddAccountInfo(rww).Build()
	}
	return ctx
}

type Account = data.Account
