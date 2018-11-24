package main

import "fmt"

func main() {
	sum(1, 2)
	sum(2, 4, 5)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)

	//printing type of functions
	fmt.Printf("sum=%T sum2=%T \n", sum, sum2)
}

//variadic function any number of arguments
func sum(nums ...int) {
	total := 0

	//looping through nums and adding it to total
	for _, num := range nums {
		total += num
	}

	fmt.Printf("nums=%v total=%d type=%T\n", nums, total, nums)
}

func sum2(nums []int) {
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Printf("nums=%v total=%d type=%T\n", nums, total, nums)
}
