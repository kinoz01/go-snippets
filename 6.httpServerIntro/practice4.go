// http.HandlerFunc vs http.HandleFunc

package main

import (
	"fmt"
	"log"
	"net/http"
)

func NightHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Good Night!")
}

func MorningHandler() http.Handler {  // You can also choose to return an http.HandlerFunc instead.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Good Morning!")
	})
}

func Router4() {
	var Night http.HandlerFunc = NightHandler
	http.Handle("/night", Night)

	http.Handle("/morning", MorningHandler())

	goodEvening := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Good Evening!")
	})
	http.HandleFunc("/evening", goodEvening)
	
	log.Println("Strating server at http://127.0.0.1:5050")
	if err := http.ListenAndServe(":5050", nil); err != nil {
		log.Fatal("Failed to start server.", err)
	}
}

