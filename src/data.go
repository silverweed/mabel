package main

import "gopkg.in/mgo.v2/bson"

type MabelConf struct {
	Title         string
	DataDir       string
	MaxUploadSize int64
	BCryptCost    int
	UserQuota     int64
}

type User struct {
	Data   UserData   `json:"data"`
	Status UserStatus `json:"status"`
}

// The user data stored in the db
type UserData struct {
	Id        bson.ObjectId `_id`
	Name      string        `json:"name"`
	Password  []byte
	Email     string
	Invite    bson.ObjectId
	MaxQuota  int64
	UsedQuota int64
}

// Volatile user status, only valid during a session
type UserStatus struct {
	Authenticated bool `json:"authenticated"`
}
