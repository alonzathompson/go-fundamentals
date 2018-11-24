package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	//Siempre Consigue Galletas - Always get the cookie
	c, err := req.Cookie("session")
	if err != nil {
		//if no session exist create one
		sID, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}

		//create cookie
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	//Set cookie
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	var u user

	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}

	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	// get cookie/ check session
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	//
	s, ok := dbSessions[c.Value]
	if ok {
		//establish time before caching s which is session struct
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}

	//set username
	_, ok = dbUsers[s.un]

	//set cookie time
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
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
