package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type reactangle struct {
	width, height float64
}

func (c circle) area() float64 {

	return 2 * math.Pi * c.radius
}

func (r reactangle) area() float64 {

	return r.height * r.width
}

func measure(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {

	c := circle{radius: 4}
	r := reactangle{4, 3}

	// first way
	fmt.Println("Circle area :", c.area())
	fmt.Println("Reactangle area :", r.area())

	//secand way
	shapes := []shape{c, r}
	fmt.Println(shapes)

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

	//third way
	//var s shape
	//fmt.Println(s.area())

	//fourth way
	measure(c)
	measure(r)

}
