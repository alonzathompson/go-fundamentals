package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	//Must establish Content-type as text/html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//Writing the html
	io.WriteString(w, `<img src="https://i.ytimg.com/vi/zlmSbJALRYQ/maxresdefault.jpg">`)
}

func main() {
	http.HandleFunc("/", d)
	http.ListenAndServe(":8080", nil)
}
