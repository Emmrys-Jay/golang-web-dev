package main

import (
	"net/http"

	"ToddMcleods/golang-web-dev/042_mongodb/10_hands-on/starting-code/controllers"
)

func main() {
	m := controllers.NewMCtrl()
	http.HandleFunc("/", m.Index)
	http.HandleFunc("/bar", m.Bar)
	http.HandleFunc("/signup", m.Signup)
	http.HandleFunc("/login", m.Login)
	http.HandleFunc("/logout", m.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
