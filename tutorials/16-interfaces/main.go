package main

import (
	"fmt"
	"math"
)

// Shape - the interface
type Shape interface {
	area() float32
	zoom(size int32)
}

type Circle struct {
	radius float32
}

type Rect struct {
	width, height float32
}

func (c *Circle) area() float32 {
	return math.Pi * c.radius
}

func (c *Circle) zoom(size int32) {
	c.radius *= float32(size)
}

func (r *Rect) area() float32 {
	return r.width * r.height
}

func (r *Rect) zoom(size int32) {
	r.width *= float32(size)
	r.height *= float32(size)
}

func main() {
	c1 := Circle{4.5}
	r1 := Rect{4, 5}

	shapes := []Shape{&c1, &r1}
	for _, s := range shapes {
		fmt.Println(s.area())
		s.zoom(2)
		fmt.Println(s.area())
	}
}
