package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/parthvinchhi/db-backup/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Config models.DbConfig
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
	return nil
}

func (m *MongoDB) RestoreMongoDBData() error {
	return nil
}
