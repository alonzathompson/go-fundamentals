package main

import (
	"fmt"
)

//AutoList to store cars and bikes
type AutoList struct {
	cars  []car
	bikes []bike
}

var autos AutoList

//car
type car struct {
	make  string
	model string
	year  string
	cyl   string
	price float32
}

//car methods
func (c car) GetInfo() string {
	//convert price to string and format float to 2 decimals
	p := fmt.Sprintf("%.2f", c.price)
	return string("make: " + c.make + " model: " + c.model + " year: " + c.year + " cyl: " + c.cyl + " price: " + p)
}

func (c car) StoreAuto() AutoList {
	autos.cars = append(autos.cars, c)
	return autos
}

//bike
type bike struct {
	make  string
	model string
	year  string
	cyl   string
	price float32
}

//bike methods
func (b bike) GetInfo() string {
	//convert price to string and format float to 2 decimals
	p := fmt.Sprintf("%.2f", b.price)
	return string("make: " + b.make + " model: " + b.model + " year: " + b.year + " cyl: " + b.cyl + " price: " + p)
}

func (b bike) StoreAuto() AutoList {
	autos.bikes = append(autos.bikes, b)
	return autos
}

//Auto interface that has the GetIfo method
//Whatever has the GetInfo method has the
//Auto interface
type Auto interface {
	GetInfo() string
	StoreAuto() AutoList
}

//Here we pass the interface into the autoInfo function
// which prints call to the GetInfo method which returns
//the info inside the struct that the GetInfo method is attached to
func autoInfo(a Auto) string {
	fmt.Println(a.GetInfo())
	return a.GetInfo()
}

func store(a Auto) AutoList {
	//fmt.Println(a.StoreAuto())
	return a.StoreAuto()
}

func main() {
	c := car{make: "lexus", model: "gs300", year: "2003", cyl: "4l", price: 2200.56}
	b := bike{make: "ducati", model: "dragon", year: "2017", cyl: "2l", price: 27000.00}

	autoInfo(c)
	autoInfo(b)
	store(c)
	store(b)

	fmt.Printf("cars: %v, bikes: %v\n", autos.cars, autos.bikes)
}
