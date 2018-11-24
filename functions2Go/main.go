package main

import "fmt"

func main() {

	// function literal
	showMsg := func(msg string) {
		fmt.Println(msg)
	}

	showMsg("Im an Object Literal")

	// IIFE
	func(msg string) {
		fmt.Println(msg)
	}("Immediately Invoked Function Expression")

	//regular func
	printMsg("Basic Bro")

}

func printMsg(msg string) {
	fmt.Println(msg)
}
