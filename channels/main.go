package main

import (
	"fmt"
	"sync"
)

//waitgroup methods Add, done,
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	//must add the amount of go routines to the wait group

	//func parameters it takes in a channel arrow pointing then chan <-chan
	//this func is recieving
	wg.Add(2)
	go func(ch <-chan int) {
		//can loop over channels
		for i := range ch {
			fmt.Println(i)
		}
		//have to notify that the wait group is done
		wg.Done()
	}(ch)

	//This func is a sender chan then <- chan<- then type
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		//must close the channel so the loop can stop on the amount of channels or will be an error
		close(ch)
		//have to notify the wait group that you are done
		wg.Done()
	}(ch)
	wg.Wait()

}
