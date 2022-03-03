package main

import (
	"ToddMcleods/golang-web-dev/042_mongodb/06_hands-on/starting-code/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/user/:id", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

/*func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}*/
