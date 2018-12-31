package main

import (
	"fmt"
	"testing"
)

func TestGetInfo(t *testing.T) {
	tables := []struct {
		a  Auto
		ex string
	}{
		{
			bike{
				make:  "suzuki",
				model: "310",
				year:  "1999",
				cyl:   "2l",
				price: 1300.00,
			},
			"make: suzuki model: 310 year: 1999 cyl: 2l price: 1300.00",
		},
		{
			bike{
				make:  "honda",
				model: "781",
				year:  "2001",
				cyl:   "2l",
				price: 1500.00,
			},
			"make: honda model: 781 year: 2001 cyl: 2l price: 1500.00",
		},
		{
			car{
				make:  "lexus",
				model: "gs300",
				year:  "1999",
				cyl:   "6l",
				price: 1700.00,
			},
			"make: lexus model: gs300 year: 1999 cyl: 6l price: 1700.00",
		},
		{
			car{
				make:  "honda",
				model: "civic",
				year:  "2001",
				cyl:   "4l",
				price: 2200.00,
			},
			"make: honda model: civic year: 2001 cyl: 4l price: 2200.00",
		},
	}

	for _, table := range tables {
		auto := autoInfo(table.a)
		if auto != table.ex {
			t.Fatalf("expected: %s got: %s", table.ex, auto)
		}
	}
}

func TestStore(t *testing.T) {
	var b AutoList
	tables := []struct {
		a  Auto
		ex string
	}{
		{
			bike{
				make:  "bmw",
				model: "450",
				year:  "1992",
				cyl:   "2l",
				price: 1100.00,
			},
			"make: bmw model: 450 year: 1992 cyl: 2l price: 1100.00",
		},
		{
			bike{
				make:  "honda",
				model: "950",
				year:  "2003",
				cyl:   "3l",
				price: 2100.00,
			},
			"make: honda model: 950 year: 2003 cyl: 3l price: 2100.00",
		},
		{
			car{
				make:  "lexus",
				model: "gs300",
				year:  "2001",
				cyl:   "6l",
				price: 2500.00,
			},
			"make: lexus model: gs300 year: 2001 cyl: 6l price: 2500.00",
		},
		{
			car{
				make:  "honda",
				model: "civic",
				year:  "2001",
				cyl:   "4l",
				price: 2200.00,
			},
			"make: honda model: civic year: 2001 cyl: 4l price: 2200.00",
		},
	}

	for _, table := range tables {
		b = store(table.a)
		fmt.Printf("autos: %v\n", table.a)
	}

	if len(b.cars) != 2 {
		t.Fatalf("expected 2 cars but got %d cars", len(b.cars))
	}

	if len(b.bikes) != 2 {
		t.Fatalf("expected 2 cars but got %d cars", len(b.bikes))
	}

}
