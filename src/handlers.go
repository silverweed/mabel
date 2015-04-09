package main

import (
	"../mustache"
	"fmt"
	"net/http"
)

func httpTemplate(template string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		send(rw, req, template, template, nil)
	}
}

func send(rw http.ResponseWriter, req *http.Request,
	name string, title string, context interface{}) {
	if len(title) > 0 {
		title = " ~ " + title
	}
	fmt.Fprintln(rw,
		mustache.RenderFileInLayout(
			mabelRoot+"/template/"+name+".html",
			mabelRoot+"/template/layout.html",
			struct {
				Title string
				Data  interface{}
			}{
				mabelConf.Title + title,
				context,
			}))
}
