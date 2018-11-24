/***********
* Templates - Basic
***********/

package main

import "fmt"

func main() {
	name := "Alonza Thompson"

	tpl := `
	<!Doctype html>
	<html lang="eng">
	<head>
	<meta charset="utf8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1> ` + name + `</h1>
	</body>
	</html>`

	fmt.Println(tpl)
}
