package main

import (
	"html/template"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

//by establishing tpl as pointer to a template we have all the methods
//available from template
var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength = 60 * 3

func init() {
	//Parse templates
	tpl = template.Must(template.ParseGlob("templates/*"))

	//generate bcrypt password - "password" is the password we pass in to encrypt
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)

	//establish a user
	dbUsers["test@test.com"] = user{"test@test.com", bs, "James", "bond", "admin"}

	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/foo", foo)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	//we are passing the w and req into our helper function
	u := getUser(w, req)

	//executing the index template and passing the data in from get user
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func foo(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	if u.Role != "admin" {
		http.Redirect(w, req, "/bar", http.StatusForbidden)
		tpl.ExecuteTemplate(w, "bar.gohtml", u)
		return
	}

	tpl.ExecuteTemplate(w, "foo.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	//if user not logged in redirect
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	//logic
	//check to see if logged in
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//check if Post has been made
	if req.Method == http.MethodPost {
		//get Form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		//check to see if username is in dbUsers[]
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "username or Password does not exist", http.StatusForbidden)

			http.Redirect(w, req, "/login", http.StatusForbidden)
			return
		}

		//does the entered password match the stored password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username or Password does not exist", http.StatusSeeOther)
		}

		//create session
		sID, err := uuid.NewV4()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		//create cookie
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		//Set Cookie
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//Execute template
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	//logic
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	c, _ := req.Cookie("session")

	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, c)

	//clean db Sessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func signup(w http.ResponseWriter, req *http.Request) {
	//logic
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//check to see if somethings been posted
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := "user"

		//check to see if username is already taken
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create session id
		sID, err := uuid.NewV4()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		//create cookie
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		//set Cookie
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}

		//encrypt password
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		//store user in db
		u := user{un, bs, f, l, r}
		//db username = user
		dbUsers[un] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//execute template
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
