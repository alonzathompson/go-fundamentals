/*****
Slices

slices can only exist on an underlying array - hence a slice of the array

the length of a slice can not be longer than the array that they come from

slices also works with strings
*****/
package main

import "fmt"

func main() {
	nums := [...]int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	sent := [...]string{"MON", "TUE", "WED", "THUR", "FRI"}
	s := "Jack and Jill went up the hill."

	// base = index 5 till the end of nums array
	base := nums[5:]
	base2 := nums[:5]

	var sum int
	var sum2 int

	for i := 0; i < len(base); i++ {
		sum += i
	}

	for _, val := range base {
		sum2 += val
	}

	//Appending to a slice
	base3 := nums[1:3]
	//extras := [...]int{21, 22, 20, 23, 24, 25}

	// pushing onto an array
	base3 = append(base3, 10, 11, 12, 13)

	//can append onto new variable, and add slice but cannot add new array
	base4 := append([]int{15, 16, 17}, base2...)

	// slice with strings
	string1 := sent[:3]
	partial := s[9:]

	/*******
	* Printing
	*******/

	fmt.Println(base, sum, sum2, base2, base3, base4)
	fmt.Printf("first 3 weekdays %s \n", string1)
	fmt.Printf("slice works with strings: %s\n", partial)
	fmt.Printf("sum1 looping index=%d \nsum1 in bytes=%08b \nsum2 total of values=%d \nsum2 in bytes=%08b\nbase values %d\nbase type=%T\nbase2=%d\nbase2 length=%d\n", sum, sum, sum2, sum2, base, base, base2, len(base2))
}
