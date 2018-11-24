/************
* Defer example two
************/

package main

import "fmt"

// Notice that the colon is not before the equal sign
var i = 0

func main() {
	fmt.Println("i(m1)=", i)
	testDefer()
	fmt.Println("i(m2)=", i)

}

func testDefer() {
	fmt.Println("i(t1)=", i)

	// The first defer function is pushed to the bottom of the stack
	defer closeFiles()
	// The second defer function is pushed next to the bottom of the stack
	// Think of defer as having it'sown stack and the first defer func call is pushed to the bottom
	// then the next one is pushed on top of that, so fourth and so on
	defer closeDBConnections()

	fmt.Println("i(t2)=", i)
	doSomething()
	fmt.Println("i(t3)=", i)
}

func closeFiles() {
	fmt.Println("i(f1)=", i)
	i = 1
}

func closeDBConnections() {
	fmt.Println("i(f2)=", i)
	i += 2
}

func doSomething() {
	fmt.Println("i(f3)=", i)
	i = 3
}
