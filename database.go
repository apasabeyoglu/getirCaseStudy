package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var ctx = context.Background()

func connectToDB() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true"))
	if err != nil {
		log.Println(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
	}
	return client, err
}

func initializeRedis() (*redis.Client, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Println(err)
	}
	return rdb, err
}
