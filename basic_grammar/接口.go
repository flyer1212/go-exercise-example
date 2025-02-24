package main

import (
	"fmt"
	"math"
)

/**

 */

type geometry interface {
	area() float64
	perim() float64
}
type rects struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rects) area() float64 {
	return r.width * r.height
}
func (r rects) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) perim3() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	r := rects{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)

}
