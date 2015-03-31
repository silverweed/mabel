package main

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

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

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	name := req.PostFormValue("name")
	if len(name) < 1 {
		http.Error(rw, "Name length < 1", http.StatusBadRequest)
		return
	}
	session.Values["name"] = name
	err := session.Save(req, rw)
	if err != nil {
		panic(err)
	}
	http.Redirect(rw, req, "/", http.StatusMovedPermanently)
}

func logoutHandler(rw http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	sessionDestroy(session, req, rw)
	http.Redirect(rw, req, "/", http.StatusMovedPermanently)
}
