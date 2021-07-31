package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

func mongoRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getFromMongoDB(w, r)
		return
	default:
		createHttpResponse(w, http.StatusMethodNotAllowed, nil, errors.New("method not allowed"))
		return
	}
}

func redisRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getFromRedis(w, r)
		return
	case "POST":
		writeToRedis(w, r)
		return
	default:
		createHttpResponse(w, http.StatusMethodNotAllowed, nil, errors.New("method not allowed"))
		return
	}
}

func getFromMongoDB(w http.ResponseWriter, r *http.Request) {
	var (
		reqBody  Request
		response Response
		recBody  RecordBody
	)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		createHttpResponse(w, http.StatusInternalServerError, nil, err)
	}

	data, err := getDataFromDB(reqBody.StartDate, reqBody.EndDate, reqBody.MinCount, reqBody.MaxCount)
	if err != nil {
		response.Code = 1
		response.Message = "failure"
		createHttpResponse(w, http.StatusInternalServerError, nil, err)
		return
	} else {
		response.Code = 0
		response.Message = "success"
		for _, k := range data {
			recBody.Key = k.ID.Key
			recBody.TotalCount = k.TotalCount
			recBody.CreatedAt = k.ID.CreatedAt
			response.Records = append(response.Records, recBody)
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		createHttpResponse(w, http.StatusInternalServerError, nil, err)
	}
	createHttpResponse(w, http.StatusOK, jsonResponse, err)
}

func getFromRedis(w http.ResponseWriter, r *http.Request) {
	var (
		request  string
		response Redis
	)

	err := r.ParseMultipartForm(0)
	if err != nil {
		createHttpResponse(w, http.StatusInternalServerError, nil, err)
	}

	request = r.FormValue("key")

	response, err = get(request)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		createHttpResponse(w, http.StatusInternalServerError, nil, err)
	}
	createHttpResponse(w, http.StatusOK, jsonResponse, err)

}

func writeToRedis(w http.ResponseWriter, r *http.Request) {
	var (
		request Redis
	)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		createHttpResponse(w, http.StatusInternalServerError, nil, err)
	}

	response, err := write(request.Key, request.Value)
	if err != nil {
		createHttpResponse(w, http.StatusInternalServerError, nil, err)
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		createHttpResponse(w, http.StatusInternalServerError, nil, err)
	}

	createHttpResponse(w, http.StatusOK, jsonResponse, err)
}

func createHttpResponse(w http.ResponseWriter, status int, jsonResponse []byte, err error) {
	if jsonResponse != nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(status)
		w.Write(jsonResponse)
	} else {
		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
	}
}
