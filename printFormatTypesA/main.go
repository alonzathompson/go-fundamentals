package main

import (
	"fmt"
)

func main() {

	s := "Hello World"
	n := 53

	//Just like C with different percent signs for displaying the data in different formats
	// binary literals are not supported in go b := 10010101
	/*
		%d = decimal
		%x = base16
	*/

	fmt.Printf("%d %#8x %#08b %T %#8o \n", n, n, n, n, n)

	n2 := 73
	fmt.Printf("%d %#8x %#08b %T %#8o \n", n2, n2, n2, n2, n2)

	fmt.Println(s)

}
