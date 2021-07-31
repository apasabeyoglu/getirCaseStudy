package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func getDataFromDB(startDate, endDate string, minCount, maxCount int32) ([]Post, error) {
	client, err := connectToDB()
	if err != nil {
		log.Println(err)
	}

	collection := client.Database("getir-case-study").Collection("records")
	sd, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Println(err)
	}

	ed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		log.Println(err)
	}

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"createdAt", bson.D{{"$lte", ed}}},
			{"createdAt", bson.D{{"$gte", sd}}},
		}}},
		{{"$project", bson.D{
			{"_id", 1},
			{"key", 1},
			{"createdAt", 1},
			{"totalCount", bson.D{{"$reduce", bson.D{{"input", "$counts"}, {"initialValue", "[ ]"}, {"in", bson.D{{"$sum", "$counts"}}}}}}}}}},
		{{"$group", bson.D{
			{"_id", bson.D{{"key", "$key"}, {"createdAt", "$createdAt"}}},
			{"totalCount", bson.D{{"$sum", "$totalCount"}}},
		}}},
		{{"$match", bson.D{
			{"totalCount", bson.D{{"$lte", maxCount}}},
			{"totalCount", bson.D{{"$gte", minCount}}},
		}}},
	}

	data, err := collection.Aggregate(ctx, pipeline)

	if err != nil {
		log.Println(err)
	}

	var result []Post
	if err = data.All(ctx, &result); err != nil {
		log.Println(err)
	}
	return result, err
}

func write(key, value string) (Redis, error) {
	rdb, err := initializeRedis()
	defer rdb.Close()

	err = rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Println(err)
	}
	return Redis{
		Key:   key,
		Value: value,
	}, err
}

func get(key string) (Redis, error) {
	rdb, err := initializeRedis()
	defer rdb.Close()

	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
		return Redis{}, err
	} else {
		return Redis{
			Key:   key,
			Value: value,
		}, nil
	}
}
