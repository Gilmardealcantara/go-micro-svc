package controllers

import (
	"net/http"

	"github.com/Gilmardealcantara/go-micro-svc/db"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/log"
)

// Hello returns a simple hello world message.
//
//	@Summary		Get hello world
//	@Description	Returns a simple hello world message.
//	@Tags			hello
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"resp: world"
//	@Router			/hello [get]
func Hello(db *db.DB, _ log.Loggable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := struct {
			Id    string `json:"id"`
			World string `json:"world"`
		}{}

		// Just to illustrate the use of the database,
		//The recommendation is to organize your code in layers using abstractions such as services, use cases, repositories...
		err := db.Pool.QueryRow(r.Context(), "select * from hello").Scan(&result.Id, &result.World)
		if err != nil {
			WriteErrorResponse(w, r, http.StatusInternalServerError, err)
			return
		}
		WriteSuccessResponse(w, result)
	}
}
