/***************
* Passing data, Variables, functions
***************/

package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// Data structure
type sage struct {
	Name, Motto string
}

//-- fm map below is an interface that holds methods to operate on our data --

// Comes with the template library - FuncMap is a key value pair of -
// key mappings to functions. Inside the html template you can activate that
// function by (using the example below) {{ uc .Name}}. uc is the function
// the . represents the data and Name is a property on that data object
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

//here we use alias tpl to cache our pointer to a template

var tpl *template.Template

func init() {
	// inside the must template, We initiate a new template then attach my
	// our Funcs method which takes an interface and we are passing our
	// funcMap interface then we parse the files.
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	// we call our Funcs function passing in fm data
	tpl = tpl.Funcs(fm)
}

// example function that is attached to fm
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	// Slice sage
	sages := []sage{
		{Name: "Steve Harvey", Motto: "Be Responsible"},
		{Name: "Ralf Waldo Emmerson", Motto: "Human Survival"},
		{Name: "Napolean Hill", Motto: "Live Raw"},
	}

	// Go pattern - to either use the err := convention to then check
	// to make sure that the err is not equal to nill
	// Or - use <var>, ok then check with an if statement to make sure
	// that it is ok
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
