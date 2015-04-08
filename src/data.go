package main

type MabelConf struct {
	Title   string
	DataDir string
}

type User struct {
	Name     string `json:"name"`
	Password []byte
	Status   UserStatus `json:"status"`
}

type UserStatus struct {
	Authenticated bool `json:"authenticated"`
}
