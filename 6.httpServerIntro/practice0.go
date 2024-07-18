// Simple server to showcase default Muliplexer usage.

package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloGet() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	}
}

func Router0() {
	defaultMux := http.NewServeMux()
	defaultMux.HandleFunc("/", HelloGet())

	log.Println("Starting a basic Hello server at http://127.0.0.1:9999")
	if err := http.ListenAndServe(":9999", defaultMux); err != nil {
		log.Fatal("Failed to start server.\n", err)
	}
}