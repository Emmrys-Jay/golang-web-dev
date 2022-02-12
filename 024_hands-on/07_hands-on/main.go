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
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("starting-files/public"))))
	http.Handle("/", http.HandlerFunc(temp))

	http.ListenAndServe(":8080", nil)
}

func temp(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template couldn't execute: ", err)
	}
}
