package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("error connecting to mongodb: ", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("error pinging mongodb: ", err)
	}

	log.Println("mongodb connected successfully")
	return client.Database("nutriflow_db")
}

func ConnectRedis() *redis.Client {
	url := os.Getenv("REDIS_URL")
	if url == "" {
		log.Fatal("REDIS_URL environment variable not set")
	}

	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal("error parsing REDIS_URL: ", err)
	}

	client := redis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal("error pinging redis: ", err)
	}

	log.Println("redis connected successfully")
	return client
}
