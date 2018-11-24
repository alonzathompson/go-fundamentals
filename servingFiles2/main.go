package main

import (
	"io"
	"net/http"
	"os"
)

func d(w http.ResponseWriter, req *http.Request) {
	//Must establish Content-type as text/html
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//Writing the html
	io.WriteString(w, `<img src="/pica.png">`)
}

func im(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("pica.png")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}

	defer f.Close()
	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", d)
	http.HandleFunc("/pica.png", im)
	http.ListenAndServe(":8080", nil)
}
