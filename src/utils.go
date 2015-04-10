package main

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// pswValidate checks if a provided (non-hashed) password
// matches a hashed value stored in the db
func pswValidate(provided string, valid []byte) bool {
	err := bcrypt.CompareHashAndPassword(valid, []byte(provided))
	return err == nil
}

// pswEncrypt generates a cryptographic hash of the given password
func pswEncrypt(password string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(password), mabelConf.BCryptCost)
	return
}

func isValidMail(email string) bool {
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
