package main

import (
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
