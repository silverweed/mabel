// Invite codes
package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type InviteCode struct {
	Id       bson.ObjectId `_id`
	Code     string
	Used     bool
	Referral bson.ObjectId
	Account  bson.ObjectId
}

func (db Database) GetInviteCode(icode string) (code *InviteCode, err error) {
	err = db.database.C("codes").Find(bson.M{"code": icode}).One(&code)
	return
}

// UseInviteCode consumes the InviteCode with Code = icode and associates it
// with a particolar User, in a one-to-one relation.
func (db Database) UseInviteCode(icode string, usedBy bson.ObjectId) error {
	// First, get the invite code and its ID
	code, err := db.GetInviteCode(icode)
	if err != nil {
		return err
	}
	// Then, update this code (mark as used, link to user account)
	op := mgo.Change{
		Update: bson.M{
			"account": usedBy,
			"used":    true,
		},
	}
	var doc InviteCode
	_, err = db.database.C("codes").FindId(code.Id).Apply(op, &doc)
	if err != nil {
		return err
	}
	// Finally, link back the user to this code
	op = mgo.Change{
		Update: bson.M{
			"invite": code.Id,
		},
	}
	var udoc UserData
	_, err = db.database.C("users").FindId(usedBy).Apply(op, &udoc)
	return err
}
