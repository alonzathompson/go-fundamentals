package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//test case
type Person struct {
	Uname string
	Email string
}

//url string
const MLABURL = "mongodb://arAdm:adminpass1@ds241658.mlab.com:41658/godata"

func handleMongo() {
	fmt.Println("Testing mlab database")
	// Do the following:
	// In a command window:
	// set MONGOLAB_URL=mongodb://IndianGuru:dbpassword@ds051523.mongolab.com:51523/godata
	// IndianGuru is my username, replace the same with yours. Type in your password.
	/*uri := os.Getenv("MLABURL")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}*/

	sess, err := mgo.Dial(MLABURL)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})

	collection := sess.DB("godata").C("supporter")

	/*err = collection.Insert(&Person{Uname: "Stefan Klaste", Email: "klaste@posteo.de"},
		&Person{Uname: "Nishant Modak", Email: "modak.nishant@gmail.com"},
		&Person{Uname: "Prathamesh Sonpatki", Email: "csonpatki@gmail.com"},
		&Person{Uname: "murtuza kutub", Email: "murtuzafirst@gmail.com"},
		&Person{Uname: "aniket joshi", Email: "joshianiket22@gmail.com"},
		&Person{Uname: "Michael de Silva", Email: "michael@mwdesilva.com"},
		&Person{Uname: "Alejandro Cespedes Vicente", Email: "cesal_vizar@hotmail.com"})
	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}*/

	result := Person{}
	err = collection.Find(bson.M{"uname": "Prathamesh Sonpatki"}).One(&result)
	if err != nil {
		log.Fatal("Error finding record: ", err)
		return
	}

	fmt.Println("Email Id:", result.Email)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	go handleMongo()
	fmt.Fprintf(w, "Checking Mlab")
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":9000", nil)
}
