package main

import "fmt"

func main() {
	//go only has for keyword for loops

	// regular for loop
	for x := 0; x < 5; x++ {
		fmt.Printf("%d", x)
	}

	// fmt.println() is the same as \n
	fmt.Println()

	// while loop style
	i := 0
	for i < 5 {
		fmt.Printf("%d", i)

		i++
	}

	fmt.Println()

	//another way to use for loops
	b := 2
	for {
		if b == 8 {
			break
		}

		if b == 4 || b == 6 {
			b++
		}

		fmt.Println(b, " ")
		b++
	}

	// multi-dimentional
	for i := 0; i <= 5; i++ {
		fmt.Printf("i = %d \n", i)
		for j := i + 1; j < 5; j++ {
			fmt.Printf("j = %d \n", j)
		}
	}

}
