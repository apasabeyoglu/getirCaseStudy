package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_GetFromMongoDB_success(t *testing.T) {
	request := Request{
		StartDate: "2001-01-01",
		EndDate:   "2021-01-01",
		MinCount:  8000,
		MaxCount:  8200,
	}
	jsonBody, err := json.Marshal(request)
	req := httptest.NewRequest("GET", "/getir/mongo", bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	getFromMongoDB(rr, req)

	require.Nil(t, err)
	require.Equal(t, 200, rr.Code)
}

func Test_GetFromRedis_success(t *testing.T) {
	err := os.Setenv("REDIS_URL", "redis://localhost:6379")
	if err != nil {
		require.NoError(t, err)
	}

	rr := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/getir", nil)

	getFromRedis(rr, req)

	require.Nil(t, err)
	require.Equal(t, 200, rr.Code)

}

func Test_WriteToRedis_success(t *testing.T) {
	err := os.Setenv("REDIS_URL", "redis://localhost:6379")
	if err != nil {
		require.Nil(t, err)
	}

	request := Redis{
		Key:   "testKey",
		Value: "testValue",
	}

	jsonBody, err := json.Marshal(request)
	req := httptest.NewRequest("POST", "/getir", bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	writeToRedis(rr, req)

	require.Nil(t, err)
	require.Equal(t, 200, rr.Code)
	require.Equal(t, "{\"key\":\"testKey\",\"value\":\"testValue\"}", rr.Body.String())
}