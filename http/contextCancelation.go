package main

import (
	"context"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/context/test-cancel", testContextCancel)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()

}

func testContextCancel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// ctx, cancel := context.WithCancel(ctx)
	// go cancelReq(cancel)
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	for true {

		select {
		case <-ctx.Done():
			println("request canceled")
			return
		case <-time.After(time.Second * 5):
			println("done request")
			return
		default:
			continue
		}
	}

}

func cancelReq(cancel context.CancelFunc) {

	time.Sleep(time.Second * 2)

	cancel()
}
