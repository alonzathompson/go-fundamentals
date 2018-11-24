package main

import (
	"html/template"
	"log"
	"net/http"
)

// you can attach the handler interface to any type
type magic int

//Here we are attaching the type hanlder interface to magic(that is what's in the reciever function)
//Type handler interface ServeHTTP(ResponseWriter, *Request)
func (m magic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "tpl.gohtml", r.Form)
}

//establishing the template
var tpl *template.Template

//initiliaze our template by parsing the template and caching it (tpl)
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	var ness magic
	// Listen and serve takes an address string and type handler
	http.ListenAndServe(":8080", ness)
}
