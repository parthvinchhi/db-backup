package db

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/parthvinchhi/db-backup/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Config models.DbConfig
	db     *gorm.DB
	Helper models.Helper
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

	log.Printf("Connected to database successfully.")

	p.db = db

	return nil
}

func (p *Postgres) BackUpPostgreSQLData() error {
	p.Helper.TimeStamp = time.Now().Format("20060102_150405")
	p.Helper.BackupFile = fmt.Sprintf("%s_backup_%s.sql", p.Config.DbName, p.Helper.TimeStamp)

	// Construct the pg_dump command
	cmd := exec.Command("pg_dump", "-U", p.Config.DbUser, "-h", p.Config.DbHost, "-F", "p", "-b", "-v", "-f", p.Helper.BackupFile, p.Config.DbName)

	// Set the environment variable for PostgreSQL password
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", p.Config.DbPassword))

	// Run the pg_dump command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("pg_dump failed: %v", err)
	}

	log.Printf("Database backup completed successfully. Backup file: %s\n", p.Helper.BackupFile)
	return nil
}

func (p *Postgres) RestorePostgreSQLData(filePath string) error {
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
	err = p.db.Exec(string(sqlContent)).Error
	if err != nil {
		log.Printf("Failed to restore the backup: %v", err)
	} else {
		log.Printf("Backup restored successfully")
	}

	return nil
}
