package main

import (
	"fmt"
)

type person struct {
	name string
	pass string
}

var people []person

func (p person) addPeople(out chan<- person) chan<- person {
	fmt.Println("GoRoutine People: Getting People")
	out <- p
	people = append(people, p)
	return out
}

func main() {
	ch := make(chan person)

	p := person{
		name: "josh",
		pass: "joshing",
	}

	go p.addPeople(ch)

	fmt.Println(<-ch, people)
}
