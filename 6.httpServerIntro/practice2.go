// Combining default server and GorillaMux server (EXAMPLE 1) running at port 3030.
// Using http.ServeMux as the Base and Adding gorilla/mux for Specific Routes.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Home2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from default ServeMux.")
}

// handler function for GorillaMux route.
func GorillaMuxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Gorilla Mux.")
}

func Router2() {
	// Create a new gorilla/mux router and define a route handled by gorilla/mux.
	gorillaRouter := mux.NewRouter()
	gorillaRouter.HandleFunc("/mux", GorillaMuxHandler).Methods("GET")

	// Create a new default http.ServeMux and define a route handled by the default http.ServeMux.
	defaultMux := http.NewServeMux()
	defaultMux.HandleFunc("/", Home2Handler)

	// Combine the routers in a new rootMux.
	rootMux := http.NewServeMux()	
	rootMux.Handle("/", defaultMux)
	rootMux.Handle("/mux", gorillaRouter)

	log.Println("Starting a server partially handled by Gorilla at http://127.0.0.1:3030")
	if err := http.ListenAndServe(":3030", rootMux); err != nil {
		log.Fatal("Failed to start server.\n", err)
	}
}
