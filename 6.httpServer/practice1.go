// Create a server the basic way at port 2020.

package main

import (
	"fmt"
	"log"
	"net/http"
)

// home page function handler.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		fmt.Fprintln(w, "Error 404 Not Found")
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusNotFound)
	}
	fmt.Fprintln(w, "Hello World!!")
}

// whoami route handler using return way.
func Whoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "My name is Ayoub, here we go again.")
	}
}

// Default way creating a server.
func Router1() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/whoami", Whoami())

	log.Println("Starting server at http://127.0.0.1:2020")
	if err := http.ListenAndServe(":2020", nil); err != nil {
		log.Fatal("Failed to start server.\n", err)
	}
}
