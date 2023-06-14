package main

import "fmt"

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}
type shape interface {
	getArea() float64
}

func main() {
	tr := triangle{
		height: 5.6,
		base:   7.8,
	}
	sq := square{
		sideLength: 9,
	}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println("The Area of", s, "is", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (t square) getArea() float64 {
	return t.sideLength * t.sideLength
}
