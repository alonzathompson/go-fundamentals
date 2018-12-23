package main

import (
	"fmt"
	"sync"
)

type cInfo struct {
	make  string
	model string
}

var cars []cInfo

var wg = sync.WaitGroup{}

func addCar(c chan<- cInfo, b cInfo) {
	c <- b
	close(c)
	wg.Done()
}

func main() {
	wg.Add(1)
	ch := make(chan cInfo)

	car := cInfo{
		make:  "Pontiac",
		model: "g6",
	}

	go addCar(ch, car)
	fmt.Println("from main println", <-ch, car)

	//fmt.Println(car, ch)
	wg.Wait()
}
