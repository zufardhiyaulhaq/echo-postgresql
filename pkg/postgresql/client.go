package postgresql

import (
	"errors"
	"fmt"
	"log"

	"github.com/zufardhiyaulhaq/echo-postgresql/pkg/settings"
	"github.com/zufardhiyaulhaq/echo-postgresql/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Interface interface {
	WriteEcho(echo *types.Echo) error
	GetEcho(id string) (types.Echo, error)
}

func New(settings settings.Settings) Client {
	if settings.PostgresqlSSLMode == "verify-full" {
		if settings.PostgresqlSSLCertPath == "" || settings.PostgresqlSSLKeyPath == "" {
			log.Fatal(errors.New("PostgresqlSSLCertPath and PostgresqlSSLKeyPath must be provided when using verify-full SSL mode"))
		}
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		settings.PostgresqlHost,
		settings.PostgresqlPort,
		settings.PostgresqlUser,
		settings.PostgresqlDatabase,
		settings.PostgresqlPassword,
		settings.PostgresqlSSLMode)

	// Add SSL certificate paths to DSN if provided
	if settings.PostgresqlSSLRootCertPath != "" {
		dsn += fmt.Sprintf(" sslrootcert=%s", settings.PostgresqlSSLRootCertPath)
	}
	if settings.PostgresqlSSLCertPath != "" {
		dsn += fmt.Sprintf(" sslcert=%s", settings.PostgresqlSSLCertPath)
	}
	if settings.PostgresqlSSLKeyPath != "" {
		dsn += fmt.Sprintf(" sslkey=%s", settings.PostgresqlSSLKeyPath)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = gormDB.AutoMigrate(&types.Echo{})
	if err != nil {
		log.Fatal(err)
	}

	return Client{
		db: gormDB,
	}
}

type Client struct {
	db *gorm.DB
}

func (s Client) WriteEcho(echo *types.Echo) error {
	tx := s.db.Save(&echo)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s Client) GetEcho(id string) (types.Echo, error) {
	var echo types.Echo

	tx := s.db.Where("echos.id = ?", id).
		First(&echo)
	if tx.Error != nil {
		return types.Echo{}, tx.Error
	}

	return echo, nil
}
