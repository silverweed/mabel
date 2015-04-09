package main

import "gopkg.in/mgo.v2/bson"

type MabelConf struct {
	Title      string
	DataDir    string
	BCryptCost int
}

type User struct {
	Data   UserData   `json:"data"`
	Status UserStatus `json:"status"`
}

type UserData struct {
	Id       bson.ObjectId `_id`
	Name     string        `json:"name"`
	Password []byte
	Email    string
}

type UserStatus struct {
	Authenticated bool `json:"authenticated"`
}
