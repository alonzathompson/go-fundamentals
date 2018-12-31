package main

import (
	"fmt"
	"testing"
)

func TestAddP(t *testing.T) {
	ch := make(chan Person)
	tables := []struct {
		x Person
		a string
	}{
		{
			Person{
				name: "A",
				pass: "one",
			},
			"A, one",
		},
		{
			Person{
				name: "B",
				pass: "two",
			},
			"B, two",
		},
		{
			Person{
				name: "C",
				pass: "three",
			},
			"C, three",
		},
	}

	for _, table := range tables {
		go table.x.AddPeople(ch)
		if <-ch != table.x {
			t.Fatalf("expected %v but got %v", table.a, <-ch)
		}
	}

	for _, tableB := range tables {
		go tableB.x.AddPeople(ch)

		var tS = fmt.Sprintf("%s, %s", tableB.x.name, tableB.x.pass)
		var temp = <-ch
		var chS = fmt.Sprintf("%s, %s", temp.name, temp.pass)

		if tS != tableB.a {
			t.Fatalf("expected: %s, and got: %s", tableB.a, chS)
		}
	}

	close(ch)

}

func TestMp(t *testing.T) {
	tables := []struct {
		a, b string
		ans  Person
	}{
		{
			a: "ro",
			b: "234",
			ans: Person{
				name: "ro",
				pass: "234",
			},
		},
		{
			a: "co",
			b: "345",
			ans: Person{
				name: "co",
				pass: "345",
			},
		},
		{
			a: "bo",
			b: "456",
			ans: Person{
				name: "bo",
				pass: "456",
			},
		},
	}

	for _, table := range tables {
		temp := MakePerson(table.a, table.b)
		if temp != table.ans {
			t.Fatalf("expected %v and got %v", table.ans, temp)
		}
	}
}
