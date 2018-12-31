package main

import (
	"testing"
)

//var wg = sync.WaitGroup{}

func TestAddCar(t *testing.T) {
	ch := make(chan CInfo)
	tables := []struct {
		x CInfo
	}{
		{
			CInfo{
				make:  "Toyota",
				model: "Camry",
			},
		},
		{
			CInfo{
				make:  "Lexus",
				model: "GS300",
			},
		},
		{
			CInfo{
				make:  "Nissan",
				model: "Maximus",
			},
		},
		{
			CInfo{
				make:  "Acura",
				model: "Nsx",
			},
		},
	}

	wg.Add(len(tables))

	for _, table := range tables {
		go SendCar(ch, table.x)
		if <-ch != table.x {
			t.Errorf("expected %+v and got %+v", table.x, <-ch)
		}

	}

	wg.Wait()
}
