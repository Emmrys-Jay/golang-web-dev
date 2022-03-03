package sessions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/satori/go.uuid"

	"ToddMcleods/golang-web-dev/042_mongodb/10_hands-on/starting-code/controllers"
	"ToddMcleods/golang-web-dev/042_mongodb/10_hands-on/starting-code/models"
)

func GetUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = controllers.SessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u models.User
	if s, ok := controllers.DbSessions[c.Value]; ok {
		s.LastActivity = time.Now()
		controllers.DbSessions[c.Value] = s
		u = controllers.DbUsers[s.Un]
	}
	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := controllers.DbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		controllers.DbSessions[c.Value] = s
	}
	_, ok = controllers.DbUsers[s.Un]
	// refresh session
	c.MaxAge = controllers.SessionLength
	http.SetCookie(w, c)
	return ok
}

func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	ShowSessions()              // for demonstration purposes
	for k, v := range controllers.DbSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(controllers.DbSessions, k)
		}
	}
	controllers.DbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	ShowSessions()             // for demonstration purposes
}

// for demonstration purposes
func ShowSessions() {
	fmt.Println("********")
	for k, v := range controllers.DbSessions {
		fmt.Println(k, v.Un)
	}
	fmt.Println("")
}
