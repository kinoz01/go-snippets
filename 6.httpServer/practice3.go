// Combining default server and GorillaMux server (EXAMPLE 2) running at port 4040.
// Using gorilla/mux as the Base and Adding Specific Routes with http.ServeMux

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router3() {
	mainRouter := mux.NewRouter()
	mainRouter.HandleFunc("/", MuxHandler).Methods("GET")

	defaultMux := http.NewServeMux()
	defaultMux.HandleFunc("/special", SpecialHandler)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/special" {
			defaultMux.ServeHTTP(w,r)
		} else {
			mainRouter.ServeHTTP(w, r)
		}
	})

	

}
