// http.Handle vs http.HandleFunc

package main

import (
	"log"
	"net/http"
)

type HeyHandler struct{}

func (hey *HeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey World!"))
}

func HeyHandleFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hey World!"))
	}
}

func HeyHandleFunc2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey World!"))
}

func Router5() {
	// Different ways to use the HeyHandler struct.
	var heyHandler HeyHandler
	http.Handle("/hey", &heyHandler)

	var heyHandler2 *HeyHandler
	http.Handle("/hey2", heyHandler2)

	heyHandler3 := HeyHandler{}
	http.Handle("/hey3", &heyHandler3)

	heyHandler4 := &HeyHandler{}
	http.Handle("/hey4", heyHandler4)

	http.Handle("/hey5", &HeyHandler{})

	http.Handle("/hey6", HeyHandleFunc())

	http.HandleFunc("/hey7", HeyHandleFunc2) // works only using HandleFunc.

	log.Println("Strating server at http://127.0.0.1:6060")
	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatal("Error starting server", err)
	}

}
