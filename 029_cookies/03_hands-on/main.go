package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var s, r int

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	s += 1
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
	fmt.Fprintln(w, "You have visited this page "+strconv.Itoa(s)+" time(s)")
}

func read(w http.ResponseWriter, req *http.Request) {
	r += 1
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
	fmt.Fprintln(w, "You have visited this page "+strconv.Itoa(r)+" time(s)")

}
