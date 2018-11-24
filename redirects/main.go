package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at foo: ", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your requst method at bar", req.Method)

	// can process form data here if you would like

	//these next two lines commented out are the same as the one http.Redirect call at the bottom
	//w.Header().Set("Location", "/")
	//w.WriteHeader(http.StatusSeeOther)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("your request method at barred: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
