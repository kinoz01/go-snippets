// Passing more complex data to the html using template.FuncMap to register functions that can be called within your template.

package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

type PageData struct {
	Names []string
	Join  func([]string, string) string
}

func Code4() {
	mux := http.NewServeMux()
	mux.HandleFunc("/c", CHandler)

	log.Println("Starting server at http://127.0.0.1:7780/c")
	if err := http.ListenAndServe(":7780", mux); err != nil {
		log.Fatal("Web server failed to start", err)
	}
}

func CHandler(w http.ResponseWriter, r *http.Request) {

	data := PageData {
		Names: GetNames(),
	}

	// Create a new template and register the 'join' function
    templ:= template.Must(template.New("CPage.html").Funcs(template.FuncMap{
        "join": strings.Join,
    }).ParseFiles("CPage.html"))

	if err := templ.Execute(w, data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
