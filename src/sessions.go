package main

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

func initSessions() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
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
