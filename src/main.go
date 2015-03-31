package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Main site configuration
var mabelConf MabelConf
var mabelRoot string

func setupHandlers(router *mux.Router) {
	GET := router.Methods("GET", "HEAD").Subrouter()
	POST := router.Methods("POST").Subrouter()

	GET.HandleFunc("/", httpHome)

	POST.HandleFunc("/login", loginHandler)
	POST.HandleFunc("/logout", logoutHandler)
}

func main() {
	// get executable path
	exec, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	mabelRoot = filepath.Dir(exec)

	// Command line parameters
	bind := flag.String("port", ":8080", "Address to bind to")
	//mongo := flag.String("dburl", "localhost", "MongoDB servers, separated by comma")
	//dbname := flag.String("dbname", "mabel", "MongoDB database to use")
	flag.StringVar(&mabelRoot, "root", mabelRoot, "The HTTP server root directory")
	flag.Parse()

	// Read configuration file
	rawconf, _ := ioutil.ReadFile("mabel.json")
	err = json.Unmarshal(rawconf, &mabelConf)
	if err != nil {
		panic(err)
	}
	initSessions()
	router := mux.NewRouter()
	setupHandlers(router)
	http.Handle("/", router)
	log.Printf("Listening on %s\r\nServer root: %s\r\n", *bind, mabelRoot)
	http.ListenAndServe(*bind, nil)
}
