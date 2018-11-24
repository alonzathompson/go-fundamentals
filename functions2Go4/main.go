/**************
* Recursion
**************/

package main

import "fmt"

func main() {

	fmt.Println(factorial(7))
	fmt.Println(factorial(5))
}

// famous factorial function
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	fmt.Print(x, " ")
	return x * factorial(x-1)
}
