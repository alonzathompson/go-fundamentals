package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am at the dog route")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am at the cat route")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Alonza Thompson")
}

func main() {
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}
