/**********
* Interfaces three - Sort package
**********/

package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{1, 4, 2, 8, 15, 7, 5}
	fmt.Println(n)

	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)

	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	// Strings

	s := []string{"Anne", "Susan", "Beatrice", "Carol", "Monique", "Regina"}
	fmt.Printf("\n%s\n", s)

	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

}
