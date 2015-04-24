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
	if !isValidMail(email) {
		http.Error(rw, "The mail you submitted is invalid", http.StatusBadRequest)
		return
	}
	// Check if invite code exists
	if _, err := db.GetInviteCode(invitecode); err != nil {
		http.Error(rw, "The invite code you used is invalid.", http.StatusTeapot)
		return
	}
	hash, err := pswEncrypt(password)
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
	if err = db.UseInviteCode(invitecode, id); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Send verification mail
	if err = users.SendRegistrationMail(email); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO
	http.Redirect(rw, req, "/", http.StatusMovedPermanently)
}

// apiSearch performs search by tag on the database.
func apiSearch(rw http.ResponseWriter, req *http.Request) {
	http.Error(rw, "Not implemented yet", http.StatusNotImplemented)
}
