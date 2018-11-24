package main

import "fmt"

func main() {

	arraysOne()
	fmt.Println()
	arraysTwo()
	fmt.Println()
	arraysThree()
}

func arraysOne() {
	var nums = [3]int{3, 4, 6}

	var sum1 int
	var sum2 int

	// len returns array/slices length
	fmt.Printf("val=%d type=%T len=%d\n", nums, nums, len(nums))

	// i is the range in nums
	for i := range nums {
		sum1 += i
		sum2 += nums[i]
	}

	fmt.Println(sum1, sum2)
}

func arraysTwo() {
	x := [...]float32{2.3, 1.5, 3.6}
	var total float32

	// for the range in x
	// <iterator>, <val> := rane in <array>
	for _, val := range x {
		total += val
	}

	fmt.Println(total)
}

func arraysThree() {
	x := [...]int{3: 10, 20}
	fmt.Println(x)
}
