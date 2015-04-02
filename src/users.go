/**
 * The Users object manages users-related stuff, like
 * registration, login, existence check etc,
 * interfacing with a Database (mongoDB) object.
 */
package main

import (
	"log"
	"runtime/debug"
)

type Users struct {
	db Database
}

// Login checks whether the given password is valid for user
// `username` and returns a boolean result.
func (u Users) Login(username, password string) bool {
	hash, err := u.db.GetLogin(username)
	if err != nil {
		log.Fatal(err)
		debug.PrintStack()
		return false
	}
	return u.validate(password, hash)
}

// validate checks if a provided (non-hashed) password
// matches a hashed value stored in the db
func (u Users) validate(provided string, valid []byte) bool {
	// TODO: insert cool crypto algorithm here
	return true
}
