package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", count)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func count(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("count")
	// if there is no cookie tha exist with the name above
	// then it will return http.ErrNoCookie
	// if ther is a ErrNoCookie then we are making sure the
	// cookie is cahced to http.Cookie with the name count and
	// value of count
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "count",
			Value: "0",
		}
	}

	// here we convert the string to an int
	count, err := strconv.Atoi(cookie.Value)

	//we add to the value
	count++

	//we then convert the int back to a string
	cookie.Value = strconv.Itoa(count)

	//we then have set the cookie
	http.SetCookie(w, cookie)

	//writing the cookie to the value
	io.WriteString(w, cookie.Value)
}
