package main

import (
	"fmt"
	"net/http"
)

type work string

func (m work) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Alonza Thompson", "Go Learner")
	res.Header().Set("Content-Type", "text/html; charset=Utf-8")
	fmt.Fprintln(res, "<h1>Any code you want in this function</h1>")
}

func main() {
	var gettem work
	http.ListenAndServe(":8080", gettem)
}
