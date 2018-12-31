package main

import "fmt"

func main() {

	fmt.Println(SumArray1([]int{3, 4, 6}))
	fmt.Println()
	fmt.Println(SumFloats([]float32{2.3, 1.5, 3.6}))
	fmt.Println()
	fmt.Println(InitValueArray([]int{3: 10, 20}))
}

func SumArray1(nums []int) int {
	var sum int

	// i is the range in nums
	for i := range nums {
		sum += nums[i]
	}

	return sum
}

func SumFloats(x []float32) float32 {
	var total float32

	for _, val := range x {
		total += val
	}

	return total
}

func InitValueArray(x []int) []int {
	return x
}
