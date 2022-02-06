package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome page!\n This is Emmanuel's Page")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I have a German Shepherd!")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "The name's Emmanuel Jonathan")
}

func main() {

	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/", index)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
