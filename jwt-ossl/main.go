package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/*To create public and private keys run the script genssl.sh
The bash command is: ./genssl.sh
This will generate your keys and make key folder*/
const (
	privKeyPath = "keys/private_key.pem"
	pubKeyPath  = "keys/public_key.pem"
)

//VerifyKey for validating tokens
var VerifyKey []byte

//SignKey for signing tokens
var SignKey []byte

//Read the Files and then assign public and private
func initKeys() {
	var err error

	SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error with private key")
		return
	}

	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error with public Key")
		return
	}
}

//GenerateJWT generates a signed json web token
func GenerateJWT() (string, error) {
	//FIRST: create a new token with a signing encryption
	token := jwt.New(jwt.SigningMethodHS256)

	//SECOND: We create the claims for the token by using
	//token(which is a pointer to a token)
	//claims (which you need a valid method, Our SinginMehodHS256)
	//the map claims puts claims in map that can be accessed as token.map
	claims := token.Claims.(jwt.MapClaims)

	//authorization
	claims["authorized"] = true
	//usr name
	claims["usr"] = "A Thompson"
	//set expiration
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	//last we establish our token string
	tokenString, err := token.SignedString(SignKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	//if there are no errors we  return our tokenstring
	return tokenString, nil
}

/*********
things to REMEMBER
--The Generate token should be used in /signup /login - or routes with data exchanges
**********/

func homePage(w http.ResponseWriter, r *http.Request) {

	//validToken recieves the return of GenerateJWT which is a token
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	c := &http.Cookie{
		Name:  "JWT",
		Value: validToken,
	}

	http.SetCookie(w, c)
	//create a new client
	client := &http.Client{}

	//Create a new http request to
	req, _ := http.NewRequest("GET", "http://localhost:9000", nil)

	//Set the header with token
	req.Header.Set("Token", validToken)

	//atatch the req to the client in the do method( the do method sends and returns http response following
	//policy such as redirects, cookies, auth )
	//Client do is returns a response from making a request
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//Read the body from the resonse with ioutil
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//Printing response from making request
	fmt.Fprintf(w, string(body))
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
