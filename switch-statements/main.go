package main

import "fmt"

func main() {
	var seasonNum int

	// Prompt String for input
	fmt.Print("Enter a Number: \n")

	// Like C language scanf takes cmd line input
	fmt.Scanf("%d", &seasonNum)

	// switch statements
	switch seasonNum {
	case 1:
		fmt.Println("spring - ", seasonNum)

	case 2:
		fmt.Println("summer - ", seasonNum)

	case 3:
		fmt.Println("fall - ", seasonNum)

	case 4:
		fmt.Println("winter - ", seasonNum)
	}
}
