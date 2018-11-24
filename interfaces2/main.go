/**************
* Interfaces - Two
*************/

package main

import "fmt"

func main() {

	var msg interface{} = "Hello"

	s, ok := msg.(string)
	if ok {
		fmt.Printf("%q %T\n", s, msg)
	} else {
		fmt.Printf("Value is not a string - %s", s)
	}
}
