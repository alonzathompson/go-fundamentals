/******
* maps
* 3 ways to declare key / value pairs
* ex1 days := make(map[string]int)
* ex2 days := map[string]int{}
*
* initialized
* ex3 days := map[string]int{
*      "Sun": 1,
*	   "Mon": 2
*	}
}
******/

package main

import "fmt"

func main() {
	d := map[string]int{
		"SUN":  0,
		"MON":  1,
		"TUE":  2,
		"WED":  3,
		"THUR": 4,
		"FRI":  5,
		"SAT":  6,
	}

	fmt.Println(d)
}
