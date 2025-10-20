package api

import (
	"net/http"

	"github.com/Gilmardealcantara/go-micro-svc/db"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/api/controllers"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(cfg *config.Configs, db *db.DB, logger log.Loggable) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /__healthcheck__", controllers.HealthCheck(db, logger))
	router.Handle("GET /swagger/", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	registerRouter(router, "GET /hello", controllers.Hello(db, logger))

	return router
}

func registerRouter(router *http.ServeMux, pattern string, finalHandler http.HandlerFunc) {
	apmMdd := tel.APMMiddleware
	logMdd := tel.Middleware(
		config.WithRequestInfoLog(true),                // success request logs are ignored by default
		config.WithHealthCheckPath("/__healthcheck__"), // default: /__healthcheck__
		//config.WithAdditionalLogHeaderData(map[string]string{"key": "X-HEADER-PROP-KEY"}),
	)
	handlerWithMiddlewares := apmMdd(logMdd(finalHandler))
	router.Handle(pattern, handlerWithMiddlewares)
}
