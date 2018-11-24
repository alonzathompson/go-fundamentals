/***********
* Structs - Anonymous fields
***********/

package main

import "fmt"

// Struct with anonymous field only accept types
// can not have multiples of the same type
type player struct {
	string
	int
}

func main() {
	p1 := player{"Muhhamed", 70}
	fmt.Println("p1 ", p1)
	fmt.Printf("p1.int=%d p1.string=%s\n", p1.int, p1.string)

	// initializing types in struct
	p2 := player{
		int:    36,
		string: "Alex Jones",
	}
	fmt.Println("p2 ", p2)
}
