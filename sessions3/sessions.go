package main

import (
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

//remember when returning to put your return type
func getUser(w http.ResponseWriter, req *http.Request) user {
	//get/check session cookie
	c, err := req.Cookie("session")

	//if no session cookie exist make one
	if err != nil {
		//got a new uuid
		sID, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}

		//make cookie - k, v pair
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	//write cookie - params writer, cookie
	http.SetCookie(w, c)

	//if user already exist get user
	var u user

	//!IMPORTANT - when getting values from map,
	// always use comma ok idiom
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

//Remember to put return type - here we are returning a bool
func alreadyLoggedIn(req *http.Request) bool {
	//we check the cookie if there is no cookie then
	//we return false user is not logged in
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	//here we get the db username if it is ok then we return a bool
	//ok idiom is checking for 0 value
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
