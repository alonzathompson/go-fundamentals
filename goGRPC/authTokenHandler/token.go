package gtoken

import (
	"fmt"
	"io/ioutil"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

/*To create public and private keys run the script genssl.sh
The bash command is: ./genssl.sh
This will generate your keys and make key folder*/
const (
	privKeyPath = "keys/keys/private_key.pem"
	pubKeyPath  = "keys/keys/public_key.pem"
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
func ParseJWT(t metadata.MD) (*jwt.Token, error) {
	var e error
	var tok *jwt.Token

	if t["token"] != nil {
		//Parse token

		token, err := jwt.Parse(t["token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}

			if token.Header["alg"] == "none" {
				return nil, fmt.Errorf("Token Header Invalid")
			}

			return VerifyKey, nil
		})

		if err != nil {
			e = err
			tok = nil
		}

		if token.Valid {
			tok = token
			e = nil
		}
	}
	return tok, e
}

/*********
things to REMEMBER
--The Generate token should be used in /signup
--Or routes that need token generation
**********/
