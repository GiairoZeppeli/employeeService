package repository

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Database string
	Port     string
	Host     string
}

func NewMongoDB(cfg Config) (*mongo.Database, error) {
	mongoURL := fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("failed to ping MongoDB: %v", err)
	}

	db := client.Database(cfg.Database)
	return db, nil
}
