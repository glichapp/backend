package main

import (
	"flag"
	"log"
	"log/slog"

	"github.com/kvizyx/glich/services/api/internal/config"
	"github.com/kvizyx/glich/services/api/internal/delivery/rest"
	sharedconfig "github.com/kvizyx/glich/shared/config"
	"github.com/kvizyx/glich/shared/logger"
)

const (
	serviceName = "api"
)

var configs sharedconfig.FlagStringSlice

func init() {
	flag.Var(
		&configs,
		"config",
		"path to the config with environment variables",
	)

	flag.Parse()
}

func main() {
	cfg, err := sharedconfig.NewFromEnv[config.Config](configs...)
	if err != nil {
		log.Fatalf("config: %s", err)
	}

	lgr, err := logger.NewSlogLogger(logger.SlogOptions{
		AppMode: cfg.App.Mode,
		Service: serviceName,
		Level:   cfg.App.LogLevel,
	})
	if err != nil {
		log.Fatalf("config: %s", err)
	}

	httpServer := rest.NewHTTPServer(cfg, lgr)

	if err = httpServer.Start(); err != nil {
		lgr.Error("failed to start http server", slog.Any("error", err))
	}
}
