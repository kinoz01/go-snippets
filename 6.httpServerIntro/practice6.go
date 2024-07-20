// Another way to start a server.

package main

import (
	"log"
	"net/http"
)

func Router6() {
	mux := http.NewServeMux()

	srv := http.Server{
		Addr: ":7070",
		Handler: mux,
	}

	log.Println("Starting server at http://127.0.0.1:7070")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Error Starting server.", err)
	}
}
