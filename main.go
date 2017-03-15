package main

import (
	"log"
	"net/http"
	"secureStore/routing"
)

func main() {
	router := routing.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}


