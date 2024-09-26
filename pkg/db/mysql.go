package db

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/parthvinchhi/db-backup/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	Config models.DbConfig
	db     *gorm.DB
	Helper models.Helper
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
	m.Helper.TimeStamp = time.Now().Format("02012006_150405")
	m.Helper.BackupFile = fmt.Sprintf("%s_backup_%s.sql", m.Config.DbName, m.Helper.TimeStamp)

	cmd := exec.Command("mysqldump", "-u", m.Config.DbUser, "-p"+m.Config.DbPassword, "-h", m.Config.DbHost, m.Config.DbName, "--result-file="+m.Helper.BackupFile)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("mysqldump failed: %v", err)
	}

	log.Printf("MySQL backup completed successfully. Backup file: %s\n", m.Helper.BackupFile)
	return nil
}

func (m *MySQL) RestoreMySQLData(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the file content using io.ReadAll
	sqlContent, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Restore the backup by executing raw SQL
	err = m.db.Exec(string(sqlContent)).Error
	if err != nil {
		log.Printf("Failed to restore the backup: %v", err)
	} else {
		log.Printf("Backup restored successfully")
	}

	return nil
}
