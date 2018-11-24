package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

	c.MaxAge = sessionLength
	//write cookie - params writer, cookie
	http.SetCookie(w, c)

	//if user already exist get user
	var u user

	//!IMPORTANT - when getting values from map,
	// always use comma ok idiom
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

//Remember to put return type - here we are returning a bool
func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	//we check the cookie if there is no cookie then
	//we return false user is not logged in
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]
	//refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok

	//here we get the db username if it is ok then we return a bool
	//ok idiom is checking for 0 value
	/*s,=un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok*/
}

func cleanSessions() {
	fmt.Println("Before Clean") //For demonstration purposes
	showSessions()              //demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second*30) || v.un == "" {
			delete(dbSessions, k)
		}
	}

	dbSessionsCleaned = time.Now()
	fmt.Println("After Clean")
	showSessions()
}

func showSessions() {
	fmt.Println("************")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
