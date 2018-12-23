package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alonzathompson/go-fundamentals/mvcpattern/models"
	"github.com/alonzathompson/go-fundamentals/mvcpattern/session"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type userLogin struct {
	userN string
	pass  string
}

//SignUp handles sign up route form data
func (c Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	if session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var u models.User
	json.NewDecoder(r.Body).Decode(&u)

	newID, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}

	if r.Method == http.MethodPost {
		//get user values
		un := u.Username
		p := u.Password
		f := u.First
		l := u.Last
		id := newID
		rl := "user"

		if _, ok := session.Users[un]; ok {
			http.Error(w, "Username already Taken", http.StatusForbidden)
			return
		}

		sID, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}

		ck := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		ck.MaxAge = session.SessionLength
		http.SetCookie(w, ck)
		session.Sessions[ck.Value] = models.Session{un, time.Now()}

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server Error", http.StatusInternalServerError)
		}

		u = models.User{un, bs, f, l, id.String(), rl}
		fmt.Println(u, "from signup")

		//username = User
		session.Users[un] = u
		fmt.Println(session.Users, "from signup sessions")

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	session.ShowSessions()
	c.tpl.ExecuteTemplate(w, "signup.gohtml", u)

}

//Login Controller for login route
func (c Controller) Login(w http.ResponseWriter, r *http.Request) {
	if session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var user userLogin

	json.NewDecoder(r.Body).Decode(&user)

	if r.Method == http.MethodPost {
		un := user.userN
		p := user.pass

		//Check for username
		u, ok := session.Users[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		//do entered match stored
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password does not match", http.StatusForbidden)
			return
		}

		//create session
		sID, _ := uuid.NewV4()
		ck := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		ck.MaxAge = session.SessionLength
		http.SetCookie(w, ck)
		session.Sessions[ck.Value] = models.Session{un, time.Now()}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

		c.tpl.ExecuteTemplate(w, "login.gohtml", nil)
	}

	session.ShowSessions()

}

//Logout controller for logout route
func (c Controller) Logout(w http.ResponseWriter, r *http.Request) {
	u := session.GetUserSession(w, r)
	fmt.Println(u)
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println(session.Sessions, "before")

	ck, _ := r.Cookie("session")
	//delete the session
	delete(session.Sessions, ck.Value)
	//remove the Cookie
	ck = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, ck)
	fmt.Println(ck, session.Sessions, "after")
	//Clean up Sessions
	if time.Now().Sub(session.LastCleaned) > (time.Second * 30) {
		go session.CleanSessions()
	}

	fmt.Println(ck, session.Sessions, "after session")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
