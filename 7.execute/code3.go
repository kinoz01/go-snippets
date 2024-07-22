// Use of Execute(w, example.html) function to response  with dynamic content.

package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Code3() {
	mux := http.NewServeMux()
	mux.HandleFunc("/a", AHandler)
	
	log.Println("Server started at http://127.0.0.1:7779/a")
	if err := http.ListenAndServe(":7779", mux); err != nil {
		log.Fatal("Error Starting server at port 7779.", err)
		return
	}
}

func AHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("APage.html")
	if err != nil {
		http.Error(w, "Could not parse template file", http.StatusInternalServerError)
		return
	}

	templ.Execute(w, GetName())
}

func GetName() (name string) {
	fmt.Print("Please enter your Name: ")
	fmt.Scanln(&name)
	return name
}
