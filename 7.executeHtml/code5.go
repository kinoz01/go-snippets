// accessing a struct methods from html using template package.

package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Pagedata struct {
	Title string
	Items []string
}

func (p Pagedata) GetTitle() string {
	return "Tile: " + p.Title
}

func (p Pagedata) GetErr() error {
	return fmt.Errorf("this is an error")
}

func Code5() {

	mux := http.NewServeMux()
	templ := template.Must(template.New("DPage.html").ParseFiles("DPage.html"))

	mux.HandleFunc("/d", func(w http.ResponseWriter, r *http.Request) {
		data := Pagedata{
			Title: "look at me mommy",
			Items: []string{"Daddy", "Bro", "Sis"},
		}
		if err := templ.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Starting server at http://127.0.0.1:7781/d")
	if err := http.ListenAndServe(":7781", mux); err != nil {
		log.Fatal("Server failed to strat", err)
	}
}
