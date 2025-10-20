package mocks

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type HttpHandler struct {
	mock.Mock
}

func (m *HttpHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	m.Called(rw, r)
}
