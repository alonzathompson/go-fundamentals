package main

import "testing"

//TestSumArray1 test for SumArray1
func TestSumArray1(t *testing.T) {
	tables := []struct {
		x []int
		y int
	}{
		{
			[]int{5, 5, 5, 5, 5},
			25,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2},
			14,
		},
		{
			[]int{3, 3, 3, 3, 3, 3, 3, 3, 3},
			27,
		},
	}

	for _, table := range tables {
		total := SumArray1(table.x)
		if total != table.y {
			t.Errorf("SumofArray1 (%d) was incorrect, got: %d, want: %d", table.x, total, table.y)
		}
	}
}

func TestFloats(t *testing.T) {
	tables := []struct {
		x []float32
		y float32
	}{
		{
			[]float32{2.3, 1.5, 3.6},
			7.3999996,
		},
		{
			[]float32{1.0, 1.5, 1.5, 1.0, 5.0},
			10.0,
		},
		{
			[]float32{3.5, 1.5, 3.5},
			8.5,
		},
	}

	for _, table := range tables {
		total := SumFloats(table.x)
		if total != table.y {
			t.Errorf("SumFloats (%f) was incorrect. got %f want: %f", table.x, total, table.y)
		}
	}
}

func TestInitArray(t *testing.T) {
	tables := []struct {
		x []int
	}{
		{
			[]int{2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{},
		},
	}

	for _, table := range tables {
		tp := InitValueArray(table.x)
		if tp == nil {
			t.Errorf("InitValueArray expected %d but got nada", table.x)
		}
	}
}
