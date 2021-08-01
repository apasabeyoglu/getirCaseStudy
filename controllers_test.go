package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_GetDataFromDB_success(t *testing.T) {
	posts, err := getDataFromDB("2001-01-01", "2021-01-01", 8000, 9000)

	require.Nil(t, err)
	require.NotNil(t, posts)
	require.Equal(t, 277, len(posts))
}

func Test_GetDataFromDB_TimeParse_error(t *testing.T) {
	posts, err := getDataFromDB("20001-001-0091", "2021-01-01", 8000, 9000)

	require.NotNil(t, err)
	require.Nil(t, posts)

	posts, err = getDataFromDB("2001-01-01", "202100-01-01", 8000, 9000)

	require.NotNil(t, err)
	require.Nil(t, posts)
}

func Test_Write_success(t *testing.T) {
	err := os.Setenv("REDIS_URL", "redis://localhost:6379")
	if err != nil {
		require.NoError(t, err)
	}
	response, err := write("testKey", "testValue")

	require.Nil(t, err)
	require.NotNil(t, response)
	require.Equal(t, &Redis{
		Key:   "testKey",
		Value: "testValue",
	}, response)
}

func Test_Write_Redis_failure(t *testing.T) {
	err := os.Setenv("REDIS_URL", "redis://localhost:9999")
	if err != nil {
		require.NoError(t, err)
	}
	response, err := write("testKey", "testValue")

	require.NotNil(t, err)
	require.Nil(t, response)

}

func Test_Get_success(t *testing.T) {
	err := os.Setenv("REDIS_URL", "redis://localhost:6379")
	if err != nil {
		require.NoError(t, err)
	}
	response, err := get("testKey")

	require.Nil(t, err)
	require.NotNil(t, response)
	require.Equal(t, &Redis{
		Key:   "testKey",
		Value: "testValue",
	}, response)
}

func Test_Get_Redis_failure(t *testing.T) {
	err := os.Setenv("REDIS_URL", "redis://localhost:9999")
	if err != nil {
		require.NoError(t, err)
	}
	response, err := get("testKey")

	require.NotNil(t, err)
	require.Nil(t, response)

}
