package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel"
)

func writeJsonResponse(w http.ResponseWriter, data any, httpCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	raw, _ := json.Marshal(data)
	_, _ = w.Write(raw)
}

func WriteSuccessResponse(w http.ResponseWriter, data any) {
	writeJsonResponse(w, data, http.StatusOK)
}

func WriteErrorResponsePayload(w http.ResponseWriter, data any, code int, err error) {
	tel.WriteHttpError(w, err)
	writeJsonResponse(w, data, code)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteErrorResponse(w http.ResponseWriter, _ *http.Request, code int, err error) {
	errResponse := ErrorResponse{
		Error: err.Error(),
	}
	if code == http.StatusInternalServerError {
		errResponse.Error = "internal server error"
	}
	WriteErrorResponsePayload(w, errResponse, code, err)
}
