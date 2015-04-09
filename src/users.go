/**
 * The Users object manages users-related stuff, like
 * registration, login, existence check etc,
 * interfacing with a Database (mongoDB) object.
 */
package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
	"runtime/debug"
	"unicode/utf8"
)

const (
	minUsernameLen = 2
	minPasswordLen = 8
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

// Register adds an user to the database, performing all due
// checks and returning a bool describing the result of the operation.
func (u Users) Register(username, password, email, invitecode string) bool {
	if utf8.RuneCountInString(username) < minUsernameLen ||
		len(password) < minPasswordLen ||
		!u.isValidMail(email) {
		return false
	}
	// Check if invite code exists
	_, err := db.GetInviteCode(invitecode)
	if err != nil {
		return false
	}
	hash, err := u.encrypt(password)
	if err != nil {
		return false
	}
	id, err := db.AddUser(username, hash, email)
	if err != nil {
		return false
	}
	// Mark invite code as used
	db.UseInviteCode(invitecode, id)
	return true
}

// validate checks if a provided (non-hashed) password
// matches a hashed value stored in the db
func (u Users) validate(provided string, valid []byte) bool {
	err := bcrypt.CompareHashAndPassword(valid, []byte(provided))
	return err == nil
}

func (u Users) encrypt(password string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(password), mabelConf.BCryptCost)
	return
}

func (u Users) isValidMail(email string) bool {
	// Translated from Javascript code of http://rosskendall.com/files/rfc822validemail.js.txt
	sQtext := `[^\x0d\x22\x5c\x80-\xff]`
	sDtext := `[^\x0d\x5b-\x5d\x80-\xff]`
	sAtom := `[^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+`
	sQuotedPair := `\x5c[\x00-\x7f]`
	sDomainLiteral := `\x5b(` + sDtext + `|` + sQuotedPair + `)*\x5d`
	sQuotedString := `\x22(` + sQtext + `|` + sQuotedPair + `)*\x22`
	sDomain_ref := sAtom
	sSubDomain := `(` + sDomain_ref + `|` + sDomainLiteral + `)`
	sWord := `(` + sAtom + `|` + sQuotedString + `)`
	sDomain := sSubDomain + `(\x2e` + sSubDomain + `)*`
	sLocalPart := sWord + `(\x2e` + sWord + `)*`
	sAddrSpec := sLocalPart + `\x40` + sDomain
	sValidEmail := `^` + sAddrSpec + `$`

	reValidEmail := regexp.MustCompile(sValidEmail)

	return reValidEmail.Match([]byte(email))
}
