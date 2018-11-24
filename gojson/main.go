package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
	Last  string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!Doctype html>
			<html lang="eng">
			<head>
			<meta charset="utf-8">
			<title>Foo</title>
			</head>
			<body>
			You are at foo
			</body>
			</html>
	`
	w.Write([]byte(s))
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Create Object
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Wry sense of humor"},
	}
	//encode with json
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
