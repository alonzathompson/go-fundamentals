package main

import "fmt"

func main() {
	s := []int{}

	s = push(s, 1, 2, 3)
	fmt.Println("push: ", s)

	s = pop(s)
	fmt.Println("pop: ", s)

	s = push(s, 10, 11)
	fmt.Println("push: ", s)

	fmt.Println("top: ", top(s))
}

// return type is a slice of int
func push(s []int, newS ...int) []int {
	return append(s, newS...)
}

// * the return type is slice of int
func pop(s []int) []int {
	return s[:len(s)-1]
}

// * the return type is just an int
func top(s []int) int {
	return s[len(s)-1]
}
