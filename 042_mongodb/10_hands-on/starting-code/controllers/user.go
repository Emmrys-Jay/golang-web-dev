package controllers

import (
	"html/template"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"ToddMcleods/golang-web-dev/042_mongodb/10_hands-on/starting-code/models"
	"ToddMcleods/golang-web-dev/042_mongodb/10_hands-on/starting-code/sessions"
)

var tpl *template.Template
var DbUsers = map[string]models.User{}       // user ID, user
var DbSessions = map[string]models.Session{} // session ID, session
var DbSessionsCleaned time.Time

const SessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*"))
	DbSessionsCleaned = time.Now()
}

type mCtrl struct{}

func NewMCtrl() (m *mCtrl) {
	return m
}

func (m *mCtrl) Index(w http.ResponseWriter, req *http.Request) {
	u := sessions.GetUser(w, req)
	sessions.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (m *mCtrl) Bar(w http.ResponseWriter, req *http.Request) {
	u := sessions.GetUser(w, req)
	if !sessions.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	sessions.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func (m *mCtrl) Signup(w http.ResponseWriter, req *http.Request) {
	if sessions.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := DbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = SessionLength
		http.SetCookie(w, c)
		DbSessions[c.Value] = models.Session{un, time.Now()}
		// store user in DbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{un, bs, f, l, r}
		DbUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	sessions.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func (m *mCtrl) Login(w http.ResponseWriter, req *http.Request) {
	if sessions.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := DbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = SessionLength
		http.SetCookie(w, c)
		DbSessions[c.Value] = models.Session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	sessions.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func (m *mCtrl) Logout(w http.ResponseWriter, req *http.Request) {
	if !sessions.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(DbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up DbSessions
	if time.Now().Sub(DbSessionsCleaned) > (time.Second * 30) {
		go sessions.CleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
