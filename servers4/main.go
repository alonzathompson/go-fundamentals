package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

//establishing a type to attach the handlr interface to
type thunder int

//Attached the ServerHTTP(w ResponseWriter, *Request) Handler Interface to type thunder
func (m thunder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//Must call ParseForm on request in order to use Form (for url and body values)
	//or PostForm (for body values)
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	//Assignment struct and initializing struct with values caching to data
	data := struct {
		Method string
		URL    *url.URL
		//Remember map[string][]string - key is type string and
		// takes a slice of values
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		//using Form - After calling ParseForm on request
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	//Execute template and pass data in
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

//initialize var tpl as a pointer to a template -
//gives us access to the template data structure
var tpl *template.Template

func init() {
	//Parses files and cache to tpl
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	//cache d as type thunder - which has the type handler interface attached to it
	var d thunder

	// Listen and serve takes type handler
	http.ListenAndServe(":8080", d)
}
