package main

type MabelConf struct {
	Title string
}

type User struct {
	Name   string     `json:"name"`
	Status UserStatus `json:"status"`
}

type UserStatus struct {
	Authenticated bool `json:"authenticated"`
}
