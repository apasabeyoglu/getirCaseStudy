package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/getir/mongo", mongoRequest)
	http.HandleFunc("/getir", redisRequest)

	// I have used this variable for both heroku and local. If you don't have port environment variable on your machine it defaults to :8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
