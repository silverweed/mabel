package main

import (
	"fmt"
	"net/http"
)

// TODO: templates
const homeHTML string = `
<!DOCTYPE html>
<html>
  <body>
    <h1>Homepage</h1>
    <form action="/login" method="POST">
      <input type="text" name="name" placeholder="name?"/>
      <input type="submit" value="auth"/>
    </form>
  </body>
</html>`

func httpHome(rw http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	// session is new? => require login
	if session.IsNew {
		fmt.Fprintf(rw, homeHTML)
	} else {
		fmt.Fprintf(rw, "You are logged in as %s\n", session.Values["name"])
	}
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	name := req.PostFormValue("name")
	if len(name) < 1 {
		panic("name.length < 1")
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
