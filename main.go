package main

import (
	"log"
	"fmt"
	"net/http"
	"secureStore/routing"
	"secureStore/config"
)

func main() {
	router := routing.NewRouter()
	addr := fmt.Sprintf(
		"%s:%s",
		config.Parameters.Server.Host,
		config.Parameters.Server.Port,
	)

	log.Fatal(http.ListenAndServe(addr, router))
}


