package token

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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
