all: mabel gulp

mabel: deps
	go build -o mabel ./src

deps:
	go get github.com/gorilla/mux
	go get github.com/gorilla/securecookie
	go get github.com/gorilla/sessions
	go get gopkg.in/mgo.v2
	go get golang.org/x/crypto/bcrypt
	touch deps

gulp:
	gulp build
	touch gulp

clean:
	rm -f deps mabel gulp

go-clean:
	rm -f mabel
