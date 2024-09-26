package db

import (
	"fmt"
	"log"

	"github.com/parthvinchhi/db-backup/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	Config models.DbConfig
	db     *gorm.DB
}

func (m *MySQL) buildConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Config.DbUser,
		m.Config.DbPassword,
		m.Config.DbHost,
		m.Config.DbPort,
		m.Config.DbName)
}

func (m *MySQL) ConnectMySQL() error {
	dsn := m.buildConnection()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	// Store the connection in the Database struct for future use
	m.db = db
	log.Println("Connected to MySQL successfully")
	return nil
}

func (m *MySQL) BackUpMySQLData() error {
	return nil
}

func (m *MySQL) RestoreMySQLData(fileName string) error {
	return nil
}
