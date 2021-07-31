package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/getir/mongo", mongoRequest)
	http.HandleFunc("/getir", redisRequest)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
