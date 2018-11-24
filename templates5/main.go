package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	tpl = tpl.Funcs(fm)
}

func dbl(x int) int {
	return x * 2
}

func sqr(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqrRt(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdbl":  dbl,
	"fsqr":  sqr,
	"fsqrt": sqrRt,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
