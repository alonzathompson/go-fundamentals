package main

import "fmt"

func main() {
	s := make([]int, 2, 3)

	a := []int{5, 6, 7, 8, 9}
	b := []int{0, 1, 2, 3, 4}

	//declaring new variable; appending initiated int slice; appending another intitiated literal slice with values
	//then appending b... as last argument; must have trailling spread on append
	c := append(make([]int, 0), append([]int{11, 12, 13}, b...)...)

	s = append(s, b...)
	b = append(b, a...)

	fmt.Println(s, a, b, c)

}
