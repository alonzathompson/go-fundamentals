package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("")
}
