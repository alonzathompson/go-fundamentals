package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am at the dog route")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am at the cat route")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "tpl.gohtml", "the data")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))
	http.Handle("/me", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
