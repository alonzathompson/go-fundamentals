package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you are at the index page")
}

func main() {
	fmt.Println("hello world\n-------------------")
	http.HandleFunc("/", index)
	http.ListenAndServe(":9000", nil)
}
