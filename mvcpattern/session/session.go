package session

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alonzathompson/go-fundamentals/mvcpattern/models"
	"github.com/satori/go.uuid"
)

const SessionLength int = 30

//var DBUsers = map[string]User{}
var Users = map[string]models.User{}
var Sessions = map[string]models.Session{}
var LastCleaned time.Time = time.Now()

func GetUserSession(w http.ResponseWriter, req *http.Request) models.User {
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
	c.MaxAge = SessionLength
	http.SetCookie(w, c)

	var u models.User

	if s, ok := Sessions[c.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
		u = Users[s.Un]
	}
	ShowSessions()
	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	// get cookie/ check session
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	//
	s, ok := Sessions[c.Value]
	if ok {
		//establish time before caching s which is session struct
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
	}

	//set username
	_, ok = Users[s.Un]

	//set cookie time
	c.MaxAge = SessionLength
	http.SetCookie(w, c)
	return ok
}

func CleanSessions() {
	fmt.Println("Before Clean") //For demonstration purposes
	ShowSessions()              //demonstration purposes
	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}

	LastCleaned = time.Now()
	fmt.Println("After Clean")
	ShowSessions()
}

func ShowSessions() {
	fmt.Println("************")
	for k, v := range Sessions {
		fmt.Println(k, v.Un)
	}
	fmt.Println("")
}
