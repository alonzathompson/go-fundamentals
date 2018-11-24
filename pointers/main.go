/*********
* Pointers
* when ever you see & (ex &x) it means address of
* ex x := 1, p := &x - p = address of x
* when ever you see * (ex *p) it means content of and can only be used with pointers
* ex x := 1, p := &x - *p = 1 because x equals 1 and p would just be address
* also &p would be the same address of x
* TIP - in go pointers are mainly used with structs
* TIP2 - a slice of an array is apointer to it's underlying array
*********/

package main

import "fmt"

func main() {
	x := 1
	pointerEx1(x)
	pointerEx2()
	pointerEx3()
	pointerEx4()
	pointerEx5()
	pointerEx6()
}

/**************************
 */

func pointerEx1(x int) {
	p := &x

	fmt.Println("Pointer basic - Example 1")
	fmt.Printf("x=%T &x=%T p=%T *p=%T &p=%T\n", x, &x, p, *p, &p)
	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x\n\n", x, &x, p, *p, &p)

	*p = 2 // derefrencing since both x and y addresses are the same,
	// changing the value of *p changes the value of x
	// because we initalized p to a pointer &x
	fmt.Printf("x=%T &x=%T p=%T *p=%T &p=%T\n", x, &x, p, *p, &p)
	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x\n\n", x, &x, p, *p, &p)

	x = 3
	fmt.Printf("x=%T &x=%T p=%T *p=%T &p=%T\n", x, &x, p, *p, &p)
	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x\n\n", x, &x, p, *p, &p)

	y := 4
	p = &y
	fmt.Printf("y=%T &y=%T p=%T *p=%T &p=%T\n\n", y, &y, p, *p, &p)

	return
}

/**************************
 */

func pointerEx2() {
	x := []int{2, 4, 6, 8}
	//reference to array value
	p := &x[0]
	p2 := &x

	fmt.Println("Pointer with Example 2")
	fmt.Printf("x[0]=%v *p=%d &x[0]=%x p=%x &p=%x x=%v p2=%v &p2=%x\n\n", x[0], *p, &x[0], p, &p, x, *p2, p2)

}

/**************************
 */

func pointerEx3() {
	x := []int{2, 4, 6, 8}
	//reference to array value
	p := &x[0]
	p2 := &x

	fmt.Println("Pointer with slice - Example 3 ")
	fmt.Printf("x[0]=%v *p=%d &x[0]=%x p=%x &p=%x x=%v p2=%v &p2=%x\n\n", x[0], *p, &x[0], p, &p, x, *p2, &p2)

}

/**************************
 */

func pointerEx4() {
	slices := [][]int{{3, 4, 6}, {2, 5, 7}, {1, 8, 9}}
	var bigSlice []int

	var p *[][]int
	p = &slices

	for r := range *p {
		bigSlice = append(bigSlice, (*p)[r]...)
	}

	fmt.Println("Pointer with slices - Example 4")
	fmt.Printf("%v %T\n", *p, *p)
	fmt.Printf("%v %T\n\n", bigSlice, bigSlice)
}

/**************************
 */

func pointerEx5() {
	var p1 = f(2)

	fmt.Println("Pointers with functions - Example 5")
	fmt.Printf("p1=%x *p1=%d &p1=%x\n", p1, *p1, &p1)

	var p2 = f(3)
	fmt.Printf("p2=%x *p2=%d &p2=%x\n\n", p2, *p2, &p2)
}

// returns a pointer thats an integer - Notice even pointers have types
func f(inp int) *int {
	v := inp * 2
	fmt.Printf("from f function &v=%x\n\n", &v)

	return &v
}

/**************************
 */
func pointerEx6() {
	a := 1

	fmt.Println("Pointer passing address to function - Example 6")
	fmt.Println()
	fmt.Printf("(1) a=%x a=%d \n", a, &a)

	// This f3 function is taking in a address and upping the value
	// by one
	f3(&a)

	// Thats why when we print it prints a 2 instead of a one because
	// we changed the value at a's address
	fmt.Printf("(2) a=%d a=%x\n", a, a)

	// passing address to function f3
	fmt.Printf("(3) a=%d a=%x\n\n", f3(&a), f3(&a))
}

// This function takes a pointer - a pointer to an int -
// it takes in a pointer location and the *int means
// we want the content at that address
func f3(y *int) int {
	//y++ increases content int by one
	*y++
	//Prints the y value (*y) and the
	fmt.Printf("(f) *y=%d y=%x \n", *y, y)
	//returns the content int - because asterisk means value
	return *y
}
