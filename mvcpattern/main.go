package main

import (
	"fmt"
	"net/http"

	"github.com/alonzathompson/go-fundamentals/mvcpattern/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := userController.NewUserController()

	r.GET("/", home)
	r.GET("/user/:id", uc.GetUser)
	r.GET("/users", uc.GetUsers)
	r.POST("/user/create", uc.CreateUser)
	r.DELETE("/user/delete/:id", uc.DeleteUser)

	http.ListenAndServe(":8080", r)
}

func home(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	s := `
	<!Doctype html>
	<html lang="en">
	<head>
	</head>
	<body>
	<h1>Welcome</h1>
	</body>
	</html>
	`

	fmt.Fprintf(w, "%s\n", s)
}
