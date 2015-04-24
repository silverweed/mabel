package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

const SESSION_NAME = "session"

var store = sessions.NewCookieStore(
	securecookie.GenerateRandomKey(32),
	securecookie.GenerateRandomKey(32))

func initSessions() {
	// Sessions' global options
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   false, // FIXME in production
	}
}

func sessionDestroy(sess *sessions.Session, req *http.Request, rw http.ResponseWriter) {
	sess.Options = &sessions.Options{
		MaxAge: -1,
		Path:   "/",
	}
	sess.Save(req, rw)
}

// login performs the actual login, i.e. sets all the required session values.
// The check on password is done by users.TryLogin, while the login HTTP request
// is processed by apiLogin.
func login(rw http.ResponseWriter, req *http.Request, session *sessions.Session, name string) error {
	session.Values["name"] = name
	return session.Save(req, rw)
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
	session, _ := store.Get(req, SESSION_NAME)
	user := User{
		Status: UserStatus{
			Authenticated: false,
		},
	}
	if !session.IsNew {
		user.Status.Authenticated = true
		user.Data.Name, _ = session.Values["name"].(string)
	}
	jsondata, err := json.Marshal(user)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(rw, string(jsondata))
}
