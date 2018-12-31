package main

import (
	"fmt"
)

type Person struct {
	name string
	pass string
}

var people []Person

func (p Person) AddPeople(out chan<- Person) chan<- Person {
	fmt.Println("GoRoutine People: Getting People")
	out <- p
	people = append(people, p)
	return out
}

func MakePerson(s1 string, s2 string) Person {
	np := Person{
		name: s1,
		pass: s2,
	}

	return np
}

func main() {
	ch := make(chan Person)
	cp := MakePerson("jon", "123")
	go cp.AddPeople(ch)

	fmt.Println(<-ch, people)
}
