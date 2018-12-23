package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "keys/private_key.pem"
	pubKeyPath  = "keys/public_key.pem"
)

//VerifyKey for validating tokens
var VerifyKey []byte

//SignKey for signing tokens
var SignKey []byte

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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information")
}

//Checks to make sure that everything is set
//This is also the correct patten to create middleware
//Creates closure enviornment of request and returns HandleFunc
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check if token is in request header
		if r.Header["Token"] != nil {
			//Parse token
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				if token.Header["alg"] == "none" {
					return nil, fmt.Errorf("Token Header Invalid")
				}

				return VerifyKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				fmt.Println(token.Header, "\n\n", token.Claims)
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequests() {
	//Homepage is protected will onlly send with token
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("my Simple jwt server")
	handleRequests()
}
