package settings

import (
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	PostgresqlEventPort string `envconfig:"PORT"`
	PostgresqlHost      string `envconfig:"POSTGRESQL_HOST"`
	PostgresqlPort      string `envconfig:"POSTGRESQL_PORT"`
	PostgresqlDatabase  string `envconfig:"POSTGRESQL_DATABASE"`
	PostgresqlUser      string `envconfig:"POSTGRESQL_USER"`
	PostgresqlPassword  string `envconfig:"POSTGRESQL_PASSWORD"`
}

func NewSettings() (Settings, error) {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		return settings, err
	}

	return settings, nil
}
