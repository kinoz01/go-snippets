package main

import (
	"fmt"
	"log"
	"net/http"
)

func Code1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
		fmt.Println("Hello") // Printed two or three times (why? ---> see doc1.md).
	})

	log.Println("Server started at http://127.0.0.1:7777")

	if err := http.ListenAndServe(":7777", nil); err != nil {
		log.Fatal("Error starting server.", err)
	}
}
