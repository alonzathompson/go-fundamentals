package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func errCh(e error) error {
	if e != nil {
		log.Fatal(e)
	}
	return e
}

func createPerson() []byte {
	elliot := &Person{
		Name: "Elliot",
		Age:  40,
	}

	data, err := proto.Marshal(elliot)
	errCh(err)
	fmt.Println("Marshalled object with protobuf", data)
	return data
}

func readPerson(b []byte) Person {
	newElliot := &Person{}
	err := proto.Unmarshal(b, newElliot)
	errCh(err)

	fmt.Println("Unmarshalled Data", newElliot)
	return *newElliot
}

func index(w http.ResponseWriter, r *http.Request) {
	d := createPerson()
	p := readPerson(d)
	fmt.Println(p)

	fmt.Fprintf(w, p.GetName())
}

func main() {
	http.HandleFunc("/", index)
	//http.HandleFunc("/main.js", handleJs)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))
}
