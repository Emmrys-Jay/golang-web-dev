package controllers

import (
	"ToddMcleods/golang-web-dev/042_mongodb/06_hands-on/starting-code/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var user = models.Users{}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// check if user is stored
	entry, ok := user[id]
	if !ok {
		fmt.Fprintln(w, "No entry with that ID available")
		return
	}

	// Marshal provided structure into JSON structure
	uj, _ := json.Marshal(entry)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//temporary storage for the decoded json
	u := models.UserVals{}

	json.NewDecoder(r.Body).Decode(&u)

	// temporarily store the user ID
	t := u.Id

	user[t] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// check if user is stored
	_, ok := user[id]
	if !ok {
		fmt.Fprintln(w, "No entry with that ID available")
		return
	}

	// Delete user
	delete(user, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}
