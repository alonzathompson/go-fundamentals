/************
* Defer Assignment
* Defer Tip - Defered expressions return after the surrounding enviornment has returned
************/

package main

import "fmt"

func main() {
	fmt.Println(square(2))
	fmt.Println(square(5))
	fmt.Println(square(6))
}

// notice that we put (result int) as return value and
// we don't have to explicitly return result at the bottom of the function
func square(x int) (result int) {
	result = x * x

	defer func() {
		if x == 2 || x == 4 {
			result += x
		}
	}()

	fmt.Println("* ")
	return
}
