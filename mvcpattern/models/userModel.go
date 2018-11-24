package models

type User struct {
	Username string //email
	Password string //password
	First    string //firstname
	Last     string //lastname
	ID       string //id
}

var Users = map[string]User{}
