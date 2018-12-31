package main

import (
	"fmt"
	"sync"
)

//CInfo -Example Car Info Object
type CInfo struct {
	make  string
	model string
}

//Cars Slice of struct Car info
var Cars []CInfo

// creating wait
var wg = sync.WaitGroup{}

//SendCar adds Car Info object (CInfo) to car channel stream
func SendCar(c chan<- CInfo, b CInfo) {
	fmt.Printf("adding car... %+v\n", b)
	c <- b
	wg.Done()
}

func main() {
	wg.Add(1)
	ch := make(chan CInfo)

	car := CInfo{
		make:  "Pontiac",
		model: "g6",
	}

	go SendCar(ch, car)
	fmt.Println("from main println", <-ch, car)

	//fmt.Println(car, ch)
	wg.Wait()
}
