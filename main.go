package main

import (
	"github.com/a-h/templ"
	"github.com/berz8/millemetri.net/templates"
	"log"
	"net/http"
)

func main() {
	static := http.FileServer(http.Dir("static"))
	index := templates.IndexPage("john")

	http.Handle("/static/", http.StripPrefix("/static/", static))
	http.Handle("/", templ.Handler(index))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
