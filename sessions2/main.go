package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

//establishing a pointer to a template
var tpl *template.Template

// remember for map map[string] = key, after that equals values inside map to key
// ex map[string]user{} = key  = map[string] and value =  user{} struct
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	//caching templates to tpl
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	//routes
	http.HandleFunc("/", home)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	//Server
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	// get cookie
	cookie, err := req.Cookie("session")
	//if cookie is not there - Create one
	if err != nil {
		sID, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
		}
		//caching cookie
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		//here we set the cookie
		http.SetCookie(w, cookie)
	}

	//if the user already exist get user
	var u user
	//the cookie.Value is the key for dbSessions map
	// notice the ok idiom patten
	// var, ok - make assignment[to key(var)] : if ok then do something
	// ok a check to see if the var in front of it has a zero value
	if un, ok := dbSessions[cookie.Value]; ok {

		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		//assigning username to un
		un := req.FormValue("username")

		//assigning firstname to f
		f := req.FormValue("firstname")

		//assigning lastname to l
		l := req.FormValue("lastname")

		//assigning username firstname lastname to u which is type user
		//this is also a composite literal
		u = user{un, f, l}

		//the value in sessions becomes the key for dbUsers
		//the session id is equal to the username
		dbSessions[cookie.Value] = un

		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	//get cookir
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		//remember to always to return
		return
	}

	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
