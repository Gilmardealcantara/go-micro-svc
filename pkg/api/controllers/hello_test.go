package controllers_test

import (
	"net/http"
	"net/http/httptest"
)

func (s *IntegrationSuite) TestHello() {
	b := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.router.ServeHTTP(b, r)

	s.Equal(http.StatusOK, b.Code)

	s.Regexp(`{"id":".*","world":"world"}`, b.Body.String())
}
