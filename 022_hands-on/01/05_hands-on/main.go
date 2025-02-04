package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "Welcome Page!")
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "I have a German Shepherd!")
	if err != nil {
		log.Fatalln(err)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "The name's Jonathan Emmanuel!")
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/dog/", http.HandlerFunc(dog))
	http.HandleFunc("/", http.HandlerFunc(index))
	http.HandleFunc("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
