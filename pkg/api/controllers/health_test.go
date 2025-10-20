package controllers_test

import (
	"net/http"
	"net/http/httptest"
)

func (s *IntegrationSuite) TestHealth() {
	b := httptest.NewRecorder()
	healthReq, _ := http.NewRequest(http.MethodGet, "/__healthcheck__", nil)
	s.router.ServeHTTP(b, healthReq)

	s.Equal(http.StatusOK, b.Code)
	s.Contains(b.Body.String(), `"status":"up"`)
	s.Contains(b.Body.String(), `"hostname":`)
}
