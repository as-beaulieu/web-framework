package main

import (
	"html/template"
	"net/http"
)

func HandleFunc(function http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		function(w, r)
	}
}

//RootHandler After receving inital "/" pattern, begin to
//Parse and send the index.html page
func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	errors(err, "Index Template Parse Error: ")
	err = tmpl.Execute(w, nil)
	errors(err, "Index Template Execution Error: ")
}
