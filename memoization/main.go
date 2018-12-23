package main

import "fmt"

/***********************
* MEMOIZATION - GO
***********************/

var memcache = make(map[int]int)

func mem() func(int) int {

	return func(n int) int {
		//checking for values in a map
		if val, ok := memcache[n]; ok {
			fmt.Println("from cache")
			return val
		} else {
			x := n + 10
			memcache[n] = x
			fmt.Println("From new")
			return x
		}
	}
}

func main() {
	newAdd := mem()
	fmt.Println(newAdd(5))
	fmt.Println(newAdd(7))
	fmt.Println(newAdd(9))
	fmt.Println(newAdd(8))
	fmt.Println(newAdd(7))
	fmt.Println(newAdd(5))
}
