package main

import (
	"fmt"
	"net/http"
	"unicode/utf8"
)

// apiSignUp is used for registering new users.
// POST params: username, password, invitecode, email
func apiSignUp(rw http.ResponseWriter, req *http.Request) {
	username := req.PostFormValue("username")
	password := req.PostFormValue("password")
	invitecode := req.PostFormValue("invitecode")
	email := req.PostFormValue("email")

	// Validate POST parameters
	if utf8.RuneCountInString(username) < minUsernameLen {
		errstr := fmt.Sprintf("Username must have at least %d characters.\n", minUsernameLen)
		http.Error(rw, errstr, http.StatusBadRequest)
		return
	}
	if len(password) < minPasswordLen {
		errstr := fmt.Sprintf("Password must have at least %d characters.\n", minPasswordLen)
		http.Error(rw, errstr, http.StatusBadRequest)
		return
	}
	if !users.IsValidMail(email) {
		http.Error(rw, "The mail you submitted is invalid", http.StatusBadRequest)
		return
	}
	// Check if invite code exists
	_, err := db.GetInviteCode(invitecode)
	if err != nil {
		http.Error(rw, "The invite code you used is invalid.", http.StatusTeapot)
		return
	}
	hash, err := users.Encrypt(password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Register user
	id, err := db.AddUser(username, hash, email)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Mark invite code as used
	err = db.UseInviteCode(invitecode, id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
