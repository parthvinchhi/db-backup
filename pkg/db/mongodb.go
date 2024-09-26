package db

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/parthvinchhi/db-backup/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Config models.DbConfig
	Helper models.Helper
	client *mongo.Client
}

func (m *MongoDB) buildConnection() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		m.Config.DbUser,
		m.Config.DbPassword,
		m.Config.DbHost,
		m.Config.DbPort,
		m.Config.DbName)
}

func (m *MongoDB) ConnectMongoDb() error {
	clientOption := options.Client().ApplyURI(m.buildConnection())

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil
	}

	m.client = client

	log.Println("Connected to MySQL successfully")

	return nil
}

func (m *MongoDB) BackUpMongoDBData() error {
	m.Helper.TimeStamp = time.Now().Format("02012006_150405")
	m.Helper.BackupFile = fmt.Sprintf("%s_backup_%s", m.Config.DbName, m.Helper.TimeStamp)

	cmd := exec.Command("mongodump", "--host", m.Config.DbHost, "--port", m.Config.DbPort, "--db", m.Config.DbName, "--out", m.Helper.BackupFile)

	// Set environment variable for MongoDB credentials
	if m.Config.DbUser != "" && m.Config.DbPassword != "" {
		cmd.Args = append(cmd.Args, "--username", m.Config.DbUser, "--password", m.Config.DbPassword)
	}

	// Run mongodump
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("mongodump failed: %v", err)
	}

	log.Printf("MongoDB backup completed successfully. Backup directory: %s\n", m.Helper.BackupFile)

	return nil
}

func (m *MongoDB) RestoreMongoDBData(backUpDir string) error {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		m.Config.DbUser,
		m.Config.DbPassword,
		m.Config.DbHost,
		m.Config.DbPort,
		m.Config.DbName)

	// Command to restore MongoDB data using mongorestore
	cmd := exec.Command("mongorestore", "--uri", uri, "--db", m.Config.DbName, "--drop", backUpDir)

	// Execute the restore command
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("MongoDB backup restored successfully")

	return nil
}
