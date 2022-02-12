package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("starting-files/templates/*"))
}

func main() {
	http.HandleFunc("/about", about)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("contact", contact)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Fatalln("template could not execute: ", err)
	}
}

func apply(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
		if err != nil {
			log.Fatalln("template could not execute: ", err)
		}
	} else if r.Method == "POST" {
		err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		if err != nil {
			log.Fatalln("template could not execute: ", err)
		}
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		log.Fatalln("template could not execute: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template could not execute: ", err)
	}
}
