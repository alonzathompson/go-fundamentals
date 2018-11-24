package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	//first we check to see if we have a cookie with seession-id
	cookie, err := req.Cookie("session-id")
	//if err != nil (if we don't have the cookie session-id)
	if err != nil {
		//we generate a new uuid
		id, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
		}
		//establish cookie args
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true,
		}
		// Notice we are setting the cookie while err is not equal to nil
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
