package main

import (
	"fmt"
	"log"
	"net/http"
)

func code1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
		fmt.Println("Hello") // Printed two or three times
	})

	log.Println("Server started at http://127.0.0.1:8888")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("Error starting server.", err)
	}
}
