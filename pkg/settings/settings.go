package settings

import (
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	HTTPPort                  string `envconfig:"HTTP_PORT" default:"80"`
	PostgresqlHost            string `envconfig:"POSTGRESQL_HOST"`
	PostgresqlPort            string `envconfig:"POSTGRESQL_PORT"`
	PostgresqlDatabase        string `envconfig:"POSTGRESQL_DATABASE"`
	PostgresqlUser            string `envconfig:"POSTGRESQL_USER"`
	PostgresqlPassword        string `envconfig:"POSTGRESQL_PASSWORD"`
	PostgresqlSSLMode         string `envconfig:"POSTGRESQL_SSL_MODE" default:"disable"`
	PostgresqlSSLRootCertPath string `envconfig:"POSTGRESQL_SSL_ROOT_CERT_PATH"`
	PostgresqlSSLCertPath     string `envconfig:"POSTGRESQL_SSL_CERT_PATH"`
	PostgresqlSSLKeyPath      string `envconfig:"POSTGRESQL_SSL_KEY_PATH"`
}

func NewSettings() (Settings, error) {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		return settings, err
	}

	return settings, nil
}
