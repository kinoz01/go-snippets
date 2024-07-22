// Use http.ServeFile to response with a static html page.
// ServeFile can only serve a static webpage.

package main

import (
	"log"
	"net/http"
)

func Code2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/static", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static.html")
	})

	log.Println("Starting server at http://127.0.0.1:7778/static")
	if err := http.ListenAndServe(":7778", mux); err != nil {
		log.Fatal("Error strating server.", err)
	}

}
