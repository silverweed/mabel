package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	if utf8.RuneCountInString(username) < MIN_USERNAME_LEN {
		errstr := fmt.Sprintf("Username must have at least %d characters.\n", MIN_USERNAME_LEN)
		http.Error(rw, errstr, http.StatusBadRequest)
		return
	}
	if len(password) < MIN_PASSWORD_LEN {
		errstr := fmt.Sprintf("Password must have at least %d characters.\n", MIN_PASSWORD_LEN)
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

// apiLogin is used to respond to authentication requests.
// The login validation itself is performed by users.Login
// in users.go.
func apiLogin(rw http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, SESSION_NAME)
	name := req.PostFormValue("name")
	password := req.PostFormValue("password")
	if len(name) < 1 {
		http.Error(rw, "Name must have at least 1 character", http.StatusBadRequest)
		return
	}
	if len(password) < 1 {
		http.Error(rw, "Empty password supplied", http.StatusBadRequest)
		return
	}
	// Validate login
	if !users.TryLogin(name, password) {
		sessionDestroy(session, req, rw)
		http.Error(rw, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	// Perform actual login
	if err := login(rw, req, session, name); err != nil {
		panic(err)
	}
	http.Redirect(rw, req, "/", http.StatusMovedPermanently)
}

func apiLogout(rw http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, SESSION_NAME)
	sessionDestroy(session, req, rw)
	http.Redirect(rw, req, "/", http.StatusMovedPermanently)
}

// apiUserData is used to retreive some session data by the client
// via AJAX, like authentication status and username.
func apiUserData(rw http.ResponseWriter, req *http.Request) {
	user, err := users.GetBySession(req)
	if err != nil {
		http.Error(rw, "Login required", http.StatusUnauthorized)
		return
	}
	jsondata, err := json.Marshal(user)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(rw, string(jsondata))
}

// apiFileUpload lets the user upload an image file to the server
func apiFileUpload(rw http.ResponseWriter, req *http.Request) {
	user, err := users.GetBySession(req)
	if err != nil {
		http.Error(rw, "Login required", http.StatusUnauthorized)
		return
	}
	// Get file from form
	file, header, err := req.FormFile("file")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Check file size
	lenstr := header.Header.Get("Content-length")
	if len(lenstr) < 1 {
		http.Error(rw, "Length required", http.StatusLengthRequired)
		return
	}
	size, err := strconv.ParseInt(lenstr, 10, 64)
	if err != nil || size > mabelConf.MaxUploadSize {
		http.Error(rw, "File size is too big.", http.StatusRequestEntityTooLarge)
		return
	}
	
	// Check user quota
	if user.Data.UsedQuota + size > user.Data.MaxQuota {
		http.Error(rw, "You reached your file upload quota. Please delete some files before submitting new ones.",
			http.StatusConflict)
		return
	}

	// Do the actual upload
	err = upload(user, file, header.Filename)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Update user quota
	db.IncQuota(user.Data.Id, size)

	// TODO: show success message
	http.Redirect(rw, req, "/", http.StatusMovedPermanently)
}
