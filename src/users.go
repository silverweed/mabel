/*
  The Users object manages users-related stuff, like
  registration, login, existence check etc,
  interfacing with a Database (mongoDB) object.
*/
package main

import (
	"log"
	"runtime/debug"
)

const (
	minUsernameLen = 2
	minPasswordLen = 8
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
