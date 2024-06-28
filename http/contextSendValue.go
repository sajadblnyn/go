package main

import (
	"context"
	"fmt"
	"net/http"
)

type ContextKeyType string

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/context/test-send-value", testSendValue)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()

}

// func testSendValue(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	ctx = context.WithValue(ctx, ContextKeyType("key"), "test")

// 	ctx = printContextValue1(ctx)

// 	printContextValue2(ctx)

// }

// func printContextValue1(ctx context.Context) context.Context {
// 	fmt.Printf("%v\n", ctx.Value(ContextKeyType("key")))

// 	ctx = context.WithValue(ctx, ContextKeyType("key"), "new_test")
// 	return ctx
// }

// func printContextValue2(ctx context.Context) {
// 	fmt.Printf("%v\n", ctx.Value(ContextKeyType("key")))

// }

func testSendValue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, ContextKeyType("key"), "test")

	printContextValue1(&ctx)

	printContextValue2(&ctx)

}

func printContextValue1(ctx *context.Context) {
	fmt.Printf("%v\n", (*ctx).Value(ContextKeyType("key")))

	(*ctx) = context.WithValue(*ctx, ContextKeyType("key"), "new_test")
}

func printContextValue2(ctx *context.Context) {
	fmt.Printf("%v\n", (*ctx).Value(ContextKeyType("key")))

}
