package main

import (
	"errors"
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
	hash = userdata.Data.Password
	return
}

// AddUser adds a new user (represented by a UserData struct) in the db,
// without doing any validation aside requiring all parameters to be non-empty.
// Returns the id of the newly created entry and possibly an error.
func (db Database) AddUser(username string, password []byte, email string) (id bson.ObjectId, err error) {
	if len(username) < 1 || len(password) < 1 || len(email) < 1 {
		err = errors.New("Empty fields in db.AddUser")
		return
	}
	id = bson.NewObjectId()
	user := UserData{
		Id:       id,
		Name:     username,
		Password: password,
		Email:    email,
	}
	err = db.database.C("users").Insert(user)
	return
}
