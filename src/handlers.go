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
		fmt.Fprintf(rw, "<!DOCTYPE html><html><body><p>You are logged in as %s</p><form action='/logout' method='POST'><input type='submit' value='logout'/></form></body></html>\n", session.Values["name"])
	}
}
