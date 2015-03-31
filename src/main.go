package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// Main site configuration
var mabelConf MabelConf

func setupHandlers(router *mux.Router) {
	GET := router.Methods("GET", "HEAD").Subrouter()
	POST := router.Methods("POST").Subrouter()

	GET.HandleFunc("/", httpHome)

	POST.HandleFunc("/login", loginHandler)
	POST.HandleFunc("/logout", logoutHandler)
}

func main() {
	// Read configuration file
	rawconf, _ := ioutil.ReadFile("mabel.json")
	err := json.Unmarshal(rawconf, &mabelConf)
	if err != nil {
		panic(err)
	}
	initSessions()
	router := mux.NewRouter()
	setupHandlers(router)
	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}
