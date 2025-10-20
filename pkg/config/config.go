package config

import (
	"fmt"
	"io"
	"os"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/data"
)

type Configs struct {
	// AppName service name
	AppName string
	// AppEnv service env name
	AppEnv string
	// BindPort port for the webserver to listen on
	BindPort string
	// BindAddress ip address of the service
	BindAddress string
	// Sentry DSN
	SentryDsn string
	// DevopsCommit commit hash of the current build
	DevopsCommit string
	// DBHost database host
	dBHost     string
	dBName     string
	dbUsername string
	dbPassword string
	dbPort     string
	dbSSLMode  string // disable
	// swagger
	SwaggerAppURL string
	// Newrelic
	NRLicenceKey   string
	NRDisabled     string
	RequestInfoLog bool
	LogLevel       data.TLogLevel
	LogOutput      io.Writer

	HealthCheckPath string

	NamespaceName string
	PodName       string
}

// not mutable at the moment
func NewConfig() Configs {
	return Configs{
		AppName:       os.Getenv("APP_NAME"),
		AppEnv:        os.Getenv("APP_ENV"),
		BindPort:      os.Getenv("BIND_PORT"),
		BindAddress:   os.Getenv("BIND_ADDRESS"),
		SentryDsn:     os.Getenv("SENTRY_DSN"),
		SwaggerAppURL: os.Getenv("SWAGGER_APP_URL"),

		NRLicenceKey:   os.Getenv("NR_LICENSE_KEY"),
		NRDisabled:     os.Getenv("NR_DISABLED"),
		RequestInfoLog: true,
		LogLevel:       os.Getenv("LOG_LEVEL"),

		dBHost:     os.Getenv("DB_HOST"),
		dBName:     os.Getenv("DB_NAME"),
		dbUsername: os.Getenv("DB_USERNAME"),
		dbPassword: os.Getenv("DB_PASSWORD"),
		dbPort:     os.Getenv("DB_PORT"),
		dbSSLMode:  os.Getenv("DB_SSL_MODE"),

		HealthCheckPath: "/__healthcheck__",

		NamespaceName: os.Getenv("NAMESPACE_NAME"),
		PodName:       os.Getenv("POD_NAME"),
		LogOutput:     os.Stdout,
	}
}

func (co Configs) DatabaseUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		co.dbUsername, co.dbPassword, co.dBHost, co.dbPort, co.dBName, co.dbSSLMode,
	)
}
