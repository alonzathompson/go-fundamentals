package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/alonzathompson/go-fundamentals/mvcpattern/session"
)

//THE CONSTRUCTOR IS A POINTER TO A TEMPLATE
type Controller struct {
	tpl *template.Template
}

// This function creates a Controller which is a pointer to a template
func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

//this reciever function has a pointer to a template as its static field
func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	u := session.GetUserSession(w, r)
	session.ShowSessions()

	//we call the template at the bottom of the function
	c.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

//Bar controller handles the Bar route
func (c Controller) Bar(w http.ResponseWriter, r *http.Request) {
	u := session.GetUserSession(w, r)
	fmt.Println(u, "user")
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	session.ShowSessions()
	c.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

//Foo Controller handles the foo route
func (c Controller) Foo(w http.ResponseWriter, r *http.Request) {
	u := session.GetUserSession(w, r)
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	if u.Role != "admin" {
		http.Redirect(w, r, "/"+u.ID+"/bar", http.StatusSeeOther)
	}

	session.ShowSessions()
	c.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
