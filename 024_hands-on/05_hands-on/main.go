package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./starting-files/templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("./starting-files/public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("Could not execute template", err)
	}
}
