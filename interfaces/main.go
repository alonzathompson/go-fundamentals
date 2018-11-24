/**********
* Interfaces
*********/

package main

import "fmt"

type rectangle struct {
	w, l int
}

func (c *rectangle) area() int {
	return c.w * c.l
}

func (c *rectangle) perim() int {
	return 2 * (c.w + c.l)
}

//+++++++++++++++++++++++

type square struct {
	s int
}

func (c *square) area() int {
	return c.s * c.s
}

func (c *square) perim() int {
	return 4 * c.s
}

//++++++++++++++++++++++++
// Interfaces are a data structure that holds methods
// because the area and the perim functions are tied to structs
// they can be used in the interface and the interface can
type shape interface {
	area() int
	perim() int
}

// The interface can take on its own methods ... notice that the
// parameter takes an interface ...The method can be applied
// to evything coming though the interface. In our example
// the area and perim methods are tied to structs. both structs
// have their own area and perim function. Those functions
// are scoped in an interface. We then add a method to the interface
// that works across both structss that are tied to to the interface
// through the methods. So when we call info, depending on the shape
// it uses the method that corresponds to the shape associated with it
func info(s shape) {
	fmt.Printf("area = %d perimeter = %d\n", s.area(), s.perim())
}

func totalArea(shapes ...shape) int {
	var totalArea int
	for _, s := range shapes {
		totalArea += s.area()
	}
	return totalArea
}

func main() {
	r1 := rectangle{2, 3}
	fmt.Printf("r1 area = %d ", r1.area())
	fmt.Printf("r1 perimeter = %d\n", r1.perim())

	r2 := rectangle{3, 7}
	fmt.Printf("r2 area = %d ", r2.area())
	fmt.Printf("r2 perimeter = %d\n", r2.perim())

	s1 := square{5}
	fmt.Printf("s1 area = %d ", s1.area())
	fmt.Printf("s1 perimeter = %d\n", s1.perim())

	s2 := square{11}
	fmt.Printf("s2 area = %d ", s2.area())
	fmt.Printf("s2 perimeter = %d\n", s2.perim())

	fmt.Println("\n==============")

	// here we are calling the info method from the shape interface on the pointer
	info(&r1)
	info(&r2)
	info(&s1)
	info(&s2)

	fmt.Println("\n===============")
	fmt.Printf("Total Area=%d ", totalArea(&r1, &r2, &s1, &s2))

	//here we make a shapes array that takes 4 elements
	// we are passing pointers to the shapes
	var shapes [4]shape
	shapes[0] = &r1
	shapes[1] = &r2
	shapes[2] = &s1
	shapes[3] = &s2

	totalArea := 0
	for _, s := range shapes {
		totalArea += s.area()
	}

	fmt.Printf("Total Area: %d", totalArea)
	info(shapes[0])
}
