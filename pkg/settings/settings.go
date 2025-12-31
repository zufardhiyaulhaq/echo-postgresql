package settings

import (
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	HTTPPort           string `envconfig:"HTTP_PORT" default:"80"`
	PostgresqlHost     string `envconfig:"POSTGRESQL_HOST"`
	PostgresqlPort     string `envconfig:"POSTGRESQL_PORT"`
	PostgresqlDatabase string `envconfig:"POSTGRESQL_DATABASE"`
	PostgresqlUser     string `envconfig:"POSTGRESQL_USER"`
	PostgresqlPassword string `envconfig:"POSTGRESQL_PASSWORD"`
	PostgresqlSSLMode  string `envconfig:"POSTGRESQL_SSL_MODE" default:"disable"`
}

func NewSettings() (Settings, error) {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		return settings, err
	}

	return settings, nil
}
