all: mabel gulp

mabel: deps
	go build -o mabel ./src

deps:
	go get github.com/gorilla/mux
	go get github.com/gorilla/securecookie
	go get github.com/gorilla/sessions
	touch deps

gulp:
	gulp build
	touch gulp

clean:
	rm -f deps mabel gulp
