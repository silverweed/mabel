// Invite codes
package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type InviteCode struct {
	Code     string
	Referral bson.ObjectId
	Account  bson.ObjectId
}

func (db Database) GetInviteCode(icode string) (code *InviteCode, err error) {
	err = db.database.C("codes").Find(bson.M{"code": icode}).One(&code)
	return
}

func (db Database) UseInviteCode(icode string, usedBy bson.ObjectId) error {
	op := mgo.Change{
		Update: bson.M{
			"account": usedBy,
		},
	}
	var doc InviteCode
	_, err := db.database.C("codes").Find(bson.M{"code": icode}).Apply(op, &doc)
	return err
}
