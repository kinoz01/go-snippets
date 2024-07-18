// Combining default server and GorillaMux server (EXAMPLE 2) running at port 4040.
// Using gorilla/mux as the Base and Adding Specific Routes with http.ServeMux

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MuxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello From GorillaMux (Default Handler)")
}

func SpecialHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello From Special page (Handled by Default http package)")
}

func Router3() {
	mainRouter := mux.NewRouter()
	mainRouter.HandleFunc("/", MuxHandler).Methods("GET")

	defaultMux := http.NewServeMux()
	defaultMux.HandleFunc("/special", SpecialHandler)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/special" {
			defaultMux.ServeHTTP(w, r)
		} else {
			mainRouter.ServeHTTP(w, r)
		}
	})

	handler2 := mux.NewRouter()
	handler2.Handle("/special", defaultMux)
	handler2.Handle("/", mainRouter)

	log.Println("Starting server at http://127.0.0.1:4040 (handler) and another at http://127.0.0.1:4041 (handler2)")

	go func() {
		if err := http.ListenAndServe(":4040", handler); err != nil {
			log.Fatal("Error starting server.", err)
		}
	}()

	go func() {
		if err := http.ListenAndServe(":4041", handler2); err != nil {
			log.Fatal("Error starting server.", err)
		}
	}()

}
