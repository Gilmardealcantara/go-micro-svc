package http

import (
	"net/http"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type nrclient struct {
	client *http.Client
	cfg    config.Configs
}

func NewHttpClientDefault(cfg config.Configs) Client {
	return &nrclient{
		client: http.DefaultClient,
	}
}

func NewHttpClient(client *http.Client, cfg config.Configs) Client {
	return &nrclient{
		client: client,
	}
}

func (nr *nrclient) Do(req *http.Request) (resp *http.Response, err error) {
	req.Header.Add("X-CLIENT-SERVICE", nr.cfg.AppName)
	txn := newrelic.FromContext(req.Context())
	s := newrelic.StartExternalSegment(txn, req)
	resp, err = nr.client.Do(req)
	s.Response = resp
	defer s.End()
	return
}
