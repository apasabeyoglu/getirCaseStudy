package main

import "time"

type Post struct {
	ID         Group `bson:"_id,omitempty"`
	TotalCount int32 `bson:"totalCount,omitempty"`
}

type Redis struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Group struct {
	Key       string    `bson:"key,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty"`
}

type Response struct {
	Code    int          `json:"code"`
	Message string       `json:"msg"`
	Records []RecordBody `json:"records"`
}

type RecordBody struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int32     `json:"totalCount"`
}

type Request struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int32  `json:"minCount"`
	MaxCount  int32  `json:"maxCount"`
}
