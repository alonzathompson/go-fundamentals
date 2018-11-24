/*********
* Structs
* Bare Minimum
*********/

package main

import "fmt"

func main() {
	type myType float64
	var total myType

	total = 44

	//notice that when we print the type of total is myType
	fmt.Printf("%.2f %T\n", total, total)

	var total2 float64
	total2 = float64(total)

	fmt.Printf("%.2f %T \n", total2, total2)
}
