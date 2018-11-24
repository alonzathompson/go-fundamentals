/**********
* Defer
* Defer - is almost like waiting for a promise to desolve
* Defer - basically is a waiting mechanism
**********/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loc 1", time.Now())

	//showMsg()
	defer showMsg()

	fmt.Println("Loc 2", time.Now())

	// function is defered for 5 seconds
	time.Sleep(5 * time.Second)
	fmt.Println("Loc 3", time.Now())
}

func showMsg() {
	fmt.Println("\nshowMsg", time.Now())
}
