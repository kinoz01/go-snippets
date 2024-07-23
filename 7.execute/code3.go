// Use of Execute(w, example.html) function to response  with dynamic content.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Code3() {

	mux := http.NewServeMux()
	mux.HandleFunc("/a", AHandler)

	mux.HandleFunc("/b", BHandler())

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

func BHandler() http.HandlerFunc {

	templ, err := template.ParseFiles("BPage.html")
	if err != nil {
		log.Fatalf("Could not parse template file: %v", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if err := templ.ExecuteTemplate(w, "BPage.html", GetNames()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}
}

func GetName() (name string) {
	fmt.Print("Please enter your Name: ")
	fmt.Scanln(&name)
	return name
}

func GetNames() []string {
	fmt.Print("Please enter names separated by a space: ")
	reader := bufio.NewReader(os.Stdin)
	phrase, _ := reader.ReadString('\n')
	return strings.Fields(phrase)
}
