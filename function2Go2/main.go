/********
* Closures
*********/

package main

import "fmt"

func main() {

	// Example one
	// Closures - Caputers first enviornment
	next := getPositiveInt()

	// this is the second function call being called three times
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	// Example two
	addCounter, multCounter := addBy(), multiBy()

	fmt.Print(addCounter(2), " ")
	fmt.Print(addCounter(3), " ")
	fmt.Print(addCounter(-1), "\n")

	fmt.Print(multCounter(2), " ")
	fmt.Print(multCounter(4), " ")
	fmt.Print(multCounter(-2), " \n")
}

// the exta func at the end means we are returning an anonymous
// func that returns an int
func getPositiveInt() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func addBy() func(int) int {
	total := 0

	return func(i int) int {
		total += i
		return total
	}
}

func multiBy() func(int) int {
	total := 1

	return func(i int) (ret int) {
		total *= i
		ret = total
		return
	}
}
