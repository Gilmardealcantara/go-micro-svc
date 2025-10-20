package config

import (
	"io"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/data"
)

func Builder(configs ...Func) Configs {
	cfg := NewConfig()
	for _, fn := range configs {
		fn(&cfg)
	}
	return cfg
}

type Func func(*Configs)

func WithVersion(version string) Func {
	return func(cfg *Configs) {
		cfg.DevopsCommit = version
	}
}

func WithEnv(appEnv string) Func {
	return func(cfg *Configs) {
		cfg.AppEnv = appEnv
	}
}

func WithServiceName(name string) Func {
	return func(cfg *Configs) {
		cfg.AppName = name
	}
}

func WithLicenceKey(licenceKey string) Func {
	return func(cfg *Configs) {
		cfg.NRLicenceKey = licenceKey
	}
}

func WithLogLevel(logLevel data.TLogLevel) Func {
	return func(cfg *Configs) {
		cfg.LogLevel = logLevel
	}
}

func WithLogOutput(logOutput io.Writer) Func {
	return func(cfg *Configs) {
		cfg.LogOutput = logOutput
	}
}

func WithHealthCheckPath(path string) Func {
	return func(cfg *Configs) {
		cfg.HealthCheckPath = path
	}
}

//func WithAdditionalLogData(key string, value any) Func {
//	return func(cfg *Configs) {
//		cfg.AdditionalLogData[key] = value
//	}
//}

func WithRequestInfoLog(enable bool) Func {
	return func(cfg *Configs) {
		cfg.RequestInfoLog = enable
	}
}
