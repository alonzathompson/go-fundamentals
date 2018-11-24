package main

import "fmt"

func main() {
	a := 5

	// Simple if/else
	if a == 2 {
		fmt.Println("a is = ", a)
	} else if a != 2 {
		fmt.Println("a is = ", a)
	}

	serverStatusOk := true

	if serverStatusOk {
		fmt.Println("server is up and running")
	}

	// Complex if/else
	if s := "FX"; serverStatusOk {
		fmt.Printf("%s server is up and running \n", s)
	}
}
