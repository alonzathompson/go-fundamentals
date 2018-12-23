package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//Secret for token
var mySigninKey = []byte("mysupersecretphrase")

//Also this is how to set up a function that returns an err
func GenerateJWT() (string, error) {
	//FIRST: create a new token with a signing encryption
	token := jwt.New(jwt.SigningMethodHS256)

	//SECOND: We create the claims for the token by using
	//token(which is a pointer to a token)
	//claims (which you need a valid method, Our SinginMehodHS256)
	claims := token.Claims.(jwt.MapClaims)

	//authorization
	claims["authorized"] = true
	//usr name
	claims["usr"] = "A Thompson"
	//set expiration
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	//last we establish our token string
	tokenString, err := token.SignedString(mySigninKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	//if there are no errors we  return our tokenstring
	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {

	//validToken recieves the return of GenerateJWT which is a token
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	//create a new client
	client := &http.Client{}

	//Create a new http request to
	req, _ := http.NewRequest("GET", "http://localhost:9000", nil)

	//Set the header with token
	req.Header.Set("Token", validToken)

	//atatch the req to the client in the do method( the do method sends and returns http response following
	//policy such as redirects, cookies, auth )
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//Read the body from the resonse with ioutil
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body), "/n", validToken, req.Header)
}

//moved the handlefunc request and server call outside of the main func
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("My simple Jwt")

	handleRequests()
}
