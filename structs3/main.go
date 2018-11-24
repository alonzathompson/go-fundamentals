/***********
* Stucts with pointers
***********/

package main

import "fmt"

type player struct {
	name, sport string
	age         int
}

func main() {
	p1 := &player{"Magic Johnson", "Basketball", 56}
	fmt.Printf("(*p1).name=%s p1.name=%s ", (*p1).name, p1.name)

	// Changing the struct property with a pointer p1 age is not 56 but 20
	(*p1).age = 20
	fmt.Println("Player1: ", (*p1))

	player2 := player{"Mike Tyson", "Boxer", 49}
	p2 := &player2

	fmt.Printf("(*p2).name=%s p2.name=%s ", (*p2).name, p2.name)
	fmt.Println("player2: ", (*p2))

}
