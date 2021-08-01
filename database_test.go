package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ConnectToDB_success(t *testing.T) {
	client, err := connectToDB()

	require.Nil(t, err)
	require.NotNil(t, client)

}

func Test_InitializeRedis_success(t *testing.T) {
	os.Setenv("REDIS_URL", "redis://localhost:6379")
	client, err := initializeRedis()

	require.Nil(t, err)
	require.NotNil(t, client)

}

func Test_InitializeRedis_failure(t *testing.T) {
	os.Setenv("REDIS_URL", "redis://localhost:9999")
	_, err := initializeRedis()

	require.NotNil(t, err)

}
