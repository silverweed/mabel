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
	"strings"
)

var (
	// Main site configuration
	mabelConf MabelConf
	// www root
	mabelRoot string
	// Database
	db Database
	// Users manager
	users Users
)

func setupHandlers(router *mux.Router) {
	GET := router.Methods("GET", "HEAD").Subrouter()
	POST := router.Methods("POST").Subrouter()

	GET.HandleFunc("/", httpHome)

	POST.HandleFunc("/login", apiLogin)
	POST.HandleFunc("/logout", apiLogout)
	POST.HandleFunc("/userdata", apiUserData)
}

func dontListDirs(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.URL.Path, "/") {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
}

func main() {
	// get executable path
	exec, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	mabelRoot = filepath.Dir(exec)

	// Command line parameters
	bind := flag.String("port", ":8000", "Address to bind to")
	mongo := flag.String("dburl", "localhost", "MongoDB servers, separated by comma")
	dbname := flag.String("dbname", "mabel", "MongoDB database to use")
	flag.StringVar(&mabelRoot, "root", mabelRoot, "The HTTP server root directory")
	flag.Parse()

	// Read configuration file
	rawconf, _ := ioutil.ReadFile("mabel.json")
	err = json.Unmarshal(rawconf, &mabelConf)
	if err != nil {
		panic(err)
	}

	initSessions()
	db = InitDatabase(*mongo, *dbname)
	users = Users{db}

	router := mux.NewRouter()
	setupHandlers(router)
	http.Handle("/", router)
	http.Handle("/static/", dontListDirs(http.StripPrefix("/static/",
		http.FileServer(http.Dir(mabelRoot+"/static")))))
	log.Printf("Listening on %s\r\nServer root: %s\r\n", *bind, mabelRoot)
	http.ListenAndServe(*bind, nil)
}
