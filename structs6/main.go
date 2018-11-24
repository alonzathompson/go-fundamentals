/**********
* Structs - adding methods - Reciever
**********/
package main

import "fmt"

type movie struct {
	name, actor string
}

// takes a variable type (m movie)
// then the name of the mehtod fullInfo()
// then return type which is string
// by placing m movie before the function
// we are associating the function with that type
func (m movie) fullInfo() string {
	return m.name + " " + m.actor
}

func main() {
	//initializing -
	m1 := movie{"The Rebellia", "Alonza Thompson"}

	fmt.Println(m1.fullInfo())

}
