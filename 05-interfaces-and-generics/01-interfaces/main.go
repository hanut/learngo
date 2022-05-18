package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Sides() uint8
	Area() int
	Perimeter() int
}

func PrintShape(s Shape) {
	// DO some stuff with the shape
	fmt.Printf("Number of Sides:\t%d\nArea:\t\t\t%d\nPerimeter:\t\t%d\n", s.Sides(), s.Area(), s.Perimeter())
}

type Square struct {
	Ls int
}

func (t Square) Sides() uint8 {
	return uint8(4)
}
func (t Square) Area() int {
	return (t.Ls * t.Ls)
}
func (t Square) Perimeter() int {
	return (t.Ls * 4)
}

type Circle struct {
	Radius int
}

func (c Circle) Sides() uint8 {
	return 0
}
func (c Circle) Area() int {
	rf := float32(c.Radius)
	return int(math.Pi * rf * rf)
}
func (c Circle) Perimeter() int {
	rf := float32(c.Radius)
	return int(2 * math.Pi * rf)
}

func main() {
	sq1 := Square{
		Ls: 4,
	}

	c1 := Circle{Radius: 4}

	PrintShape(sq1)
	PrintShape(c1)

}
