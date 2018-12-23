package models

import "time"

type User struct {
	Username string //email
	Password []byte //password
	First    string //firstname
	Last     string //lastname
	ID       string //id
	Role     string //user role
}

type Session struct {
	Un           string
	LastActivity time.Time
}
