package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Database struct {
	session  *mgo.Session
	database *mgo.Database
}

func InitDatabase(servers, dbname string) Database {
	var db Database
	var err error
	db.session, err = mgo.Dial(servers)
	if err != nil {
		panic(err)
	}
	db.database = db.session.DB(dbname)
	return db
}

func (db Database) Close() {
	db.session.Close()
}

// GetLogin retreives the login password hash for the user `username`.
func (db Database) GetLogin(username string) (hash []byte, err error) {
	var userdata User
	err = db.database.C("users").Find(bson.M{"name": username}).One(&userdata)
	if err != nil {
		return
	}
	hash = userdata.Password
	return
}
