package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppMode string

const (
	AppModeDevelopment AppMode = "development"
	AppModeProduction  AppMode = "production"
)

type Shared struct {
	Postgres Postgres
}

type Postgres struct {
	URL string `env:"POSTGRES_URL" env-required:"true"`
}

// NewFromEnv read files with environment variables by provided paths
// and load them along with already loaded environment variables.
func NewFromEnv[Config any](paths ...string) (Config, error) {
	var config Config

	for _, path := range paths {
		if len(path) == 0 {
			continue
		}

		if err := godotenv.Load(path); err != nil {
			return config, fmt.Errorf("load environment variables from %s: %w", path, err)
		}
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		return config, fmt.Errorf("read environment variables into config: %w", err)
	}

	return config, nil
}
