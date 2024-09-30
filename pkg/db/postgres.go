package db

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/parthvinchhi/db-backup/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresDB interface {
	ConnectPostgreSQL(ctx context.Context) error
	BackUpPostgreSQLData(ctx context.Context) error
	RestorePostgreSQLData(ctx context.Context, filePath string) error
}

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

func (p *Postgres) ConnectPostgreSQL(ctx context.Context) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  p.buildConnection(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	// Use a context for controlling connection timeouts and cancellations
	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Printf("Connected to database successfully.")

	p.db = db

	return nil
}

func (p *Postgres) BackUpPostgreSQLData(ctx context.Context) error {
	p.Helper.TimeStamp = time.Now().Format("02012006_150405")
	p.Helper.BackupFile = fmt.Sprintf("%s_backup_%s.sql", p.Config.DbName, p.Helper.TimeStamp)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	// Construct the pg_dump command
	// cmd := exec.Command("pg_dump", "-U", p.Config.DbUser, "-h", p.Config.DbHost, "-F", "p", "-b", "-v", "-f", p.Helper.BackupFile, p.Config.DbName)
	cmd := exec.CommandContext(ctx, "pg_dump", "-U", p.Config.DbUser, "-h", p.Config.DbHost, "-F", "p", "-b", "-v", "-f", p.Helper.BackupFile, p.Config.DbName)

	// Set the environment variable for PostgreSQL password
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", p.Config.DbPassword))

	cmd.Stderr = os.Stderr

	// Run the pg_dump command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("pg_dump failed: %v", err)
	}

	log.Printf("Database backup completed successfully. Backup file: %s\n", p.Helper.BackupFile)
	return nil
}

func (p *Postgres) RestorePostgreSQLData(ctx context.Context, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open backup file: %w", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Start a transaction for the restore process
	tx := p.db.Begin()

	for scanner.Scan() {
		sqlStatement := scanner.Text()

		// Skip empty lines and comments
		if len(sqlStatement) == 0 || sqlStatement[0] == '-' {
			continue
		}

		// Execute the SQL statement using GORM
		if err := tx.Exec(sqlStatement).Error; err != nil {
			// Rollback if any error occurs
			tx.Rollback()
			return fmt.Errorf("failed to execute statement: %s, error: %w", sqlStatement, err)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to read backup file: %w", err)
	}

	// Commit the transaction after successful execution of all statements
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Println("Data restored successfully from the backup file.")
	return nil
}
