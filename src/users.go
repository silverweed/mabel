/*
  The Users object manages users-related stuff, like
  registration, login, existence check etc,
  interfacing with a Database (mongoDB) object.
*/
package main

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"
)

const (
	MIN_USERNAME_LEN = 2
	MIN_PASSWORD_LEN = 8
)

type Users struct {
	db Database
}

// TryLogin checks whether the given password is valid for user
// `username` and returns a boolean result.
func (u Users) TryLogin(username, password string) bool {
	hash, err := u.db.GetLogin(username)
	if err != nil {
		log.Fatal(err)
		debug.PrintStack()
		return false
	}
	return pswValidate(password, hash)
}

func (u Users) SendRegistrationMail(email string) error {
	// TODO
	return nil
}

// GetBySession is used to retreive user data based on the client's
// session token. Will return error != nil if the session token is
// invalid (i.e. the user isn't logged in)
func (u Users) GetBySession(req *http.Request) (user User, err error) {
	session, err := store.Get(req, SESSION_NAME)
	if err != nil {
		// session couldn't be decoded
		return
	}
	user.Status.Authenticated = false
	if !session.IsNew {
		user.Status.Authenticated = true
		user.Data.Name, _ = session.Values["name"].(string)
		return
	}
	err = errors.New("User is not authenticated")
	return
}
