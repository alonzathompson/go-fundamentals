/*********
* CallBacks
* TIP - Callbacks are valuable but can increase complexity
* Only use CallBacks when you have to
*********/
package main

import "fmt"

func main() {
	square := func(i int) int {
		return i * i
	}

	cube := func(i int) int {
		return i * i * i
	}

	fmt.Printf("%v %b \n", calc(square, 3), calc(square, 3))
	fmt.Printf("%v %b \n", calc(cube, 8), calc(square, 3))

	// Callback function javascript style with a anonymous function as the last argument
	fmt.Printf("%v\n", calc(func(i int) int {
		return i * i
	}, 3))
}

// function that takes a function as a parameter and another interger
// A callback function
func calc(f func(int) int, x int) int {
	return f(x)
}
