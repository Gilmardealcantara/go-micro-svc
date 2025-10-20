package controllers

import (
	"net/http"
	"os"

	"github.com/Gilmardealcantara/go-micro-svc/db"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/log"
)

// HealthCheck returns the current health status of the application.
//
//	@Summary		Get healthcheck
//	@Description	Returns the current health status of the application.
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	healthPayload
//	@Failure		500	{object}	healthPayload
//	@Router			/__healthcheck__ [get]
func HealthCheck(db *db.DB, _ log.Loggable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		err := db.Pool.Ping(r.Context())
		if err != nil {
			WriteErrorResponsePayload(w, healthPayload{Hostname: hostname, Status: "down"}, http.StatusInternalServerError, err)
			return
		}
		WriteSuccessResponse(w, healthPayload{Hostname: hostname, Status: "up"})
	}
}

type healthPayload struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}
