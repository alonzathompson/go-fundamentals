package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You have arrived at dog central")
}

func c(w http.ResponseWriter, req *http.Request) {
	//the io can write to the browser or the terminal
	// here we are writing to the response
	io.WriteString(w, "You have arrived at catville lane")
}

func main() {

	//the handler func takes two parameters -
	// the pattern string, and a func with type handler
	// type handler is implemented by attaching ResponseWriter
	// and pointer to a Request (w http.ResponseWriter, req *http.Request)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	//By passing nil to listen and serve we use the
	//default server mux this allows us to implement the HandleFunc
	http.ListenAndServe(":8080", nil)
}
