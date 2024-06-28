package main

import (
	"net/http"
	"time"
)

func main() {

	// http.ListenAndServe(":8080", nil)
	server := &http.Server{
		Addr:         ":9090",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
