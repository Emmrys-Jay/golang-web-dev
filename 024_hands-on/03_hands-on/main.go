package main

import "net/http"

func main() {
	web := http.FileServer(http.Dir("./starting-files/"))

	http.Handle("/", web)
	http.ListenAndServe(":8080", nil)
}
