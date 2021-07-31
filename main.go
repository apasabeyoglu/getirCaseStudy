package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/getir/mongo", mongoRequest)
	http.HandleFunc("/getir", redisRequest)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
