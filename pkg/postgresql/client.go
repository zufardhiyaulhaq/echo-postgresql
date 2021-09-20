package postgresql

import (
	"fmt"
	"log"

	"github.com/zufardhiyaulhaq/echo-postgresql/pkg/settings"
	"github.com/zufardhiyaulhaq/echo-postgresql/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresqlClient struct {
	db *gorm.DB
}

func New(settings settings.Settings) PostgresqlClient {
	dsn := fmt.Sprintf("host=%s port=%s user=%s DB.name=%s password=%s sslmode=disable",
		settings.PostgresqlHost,
		settings.PostgresqlPort,
		settings.PostgresqlUser,
		settings.PostgresqlDatabase,
		settings.PostgresqlPassword)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = gormDB.AutoMigrate(&types.Echo{})
	if err != nil {
		log.Fatal(err)
	}

	return PostgresqlClient{
		db: gormDB,
	}
}

func (s *PostgresqlClient) WriteEcho(echo *types.Echo) error {
	tx := s.db.Save(&echo)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *PostgresqlClient) GetEcho(id string) (types.Echo, error) {
	var echo types.Echo

	tx := s.db.Where("echo.id = ?", id).
		First(&echo)
	if tx.Error != nil {
		return types.Echo{}, tx.Error
	}

	return echo, nil
}
