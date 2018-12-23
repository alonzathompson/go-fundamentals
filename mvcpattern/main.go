package main

import (
	"net/http"

	"html/template"

	"github.com/alonzathompson/go-fundamentals/mvcpattern/controllers"
	//"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	//r := httprouter.New()
	c := controllers.NewController(tpl)

	http.HandleFunc("/", c.Index)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/signup", c.SignUp)
	http.HandleFunc("/logout", c.Logout)
	http.HandleFunc("/bar", c.Bar)
	http.HandleFunc("/foo", c.Foo)

	//http.Handle("/favicon.ico", http.NotFoundHandler())
	//http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
