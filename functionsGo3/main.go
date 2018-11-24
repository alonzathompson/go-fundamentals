package main

import "fmt"

func main() {
	data := []string{"Daisy", "Rose", "Tulip"}

	fmt.Printf("%q\n", trimSlice(data))

	fmt.Printf("%q\n", data)
}

func trimSlice(data []string) []string {
	var newData []string

	for _, d := range data {
		if d != "" {
			newData = append(newData, d)
		}
	}

	return newData
}
