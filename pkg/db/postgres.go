package db

import (
	"fmt"

	"github.com/parthvinchhi/db-backup/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Config models.DbConfig
	db     *gorm.DB
}

func (p *Postgres) buildConnection() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Config.DbHost,
		p.Config.DbPort,
		p.Config.DbUser,
		p.Config.DbPassword,
		p.Config.DbName,
		p.Config.DbSslMode)
}

func (p *Postgres) ConnectPostgreSQL() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  p.buildConnection(),
		PreferSimpleProtocol: true,
	}))

	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err == nil {
		err = sqlDB.Ping()
		if err != nil {
			return err
		}
	}

	p.db = db

	return nil
}
