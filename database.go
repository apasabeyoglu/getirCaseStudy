package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func connectToDB() (*mongo.Client, error) {
	// Normally this URI should be hidden and used Hashicorp Vault or something like that for safety reasons. However I've wanted to avoid over engineering.
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
	// For local testing make sure that you have set environment variable for "REDIS_URL". I have implemented a solution in the tests those need to use Redis.
	url, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(url)

	// I wanted to be sure that db is up and running so I send ping to redis
	err = rdb.Ping(ctx).Err()
	if err != nil {
		log.Println(err)
	}
	return rdb, err
}
