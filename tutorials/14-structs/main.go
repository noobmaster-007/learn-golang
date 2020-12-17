package main

import "fmt"

// Point create a struct
// type <name> struct
type Point struct {
	x int32
	y int32
}

// Circle nested struct
// name of struct pointer is optional
// it can be access directly. e.g. c1.x, c1.y
type Circle struct {
	radius float64
	*Point
}

// when passing pointer for struct
// there is no need to de reference
func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	p1 := Point{1, 2}
	fmt.Println(p1)

	p2 := &Point{x: 2}
	fmt.Println(p2)
	changeX(p2)
	fmt.Println(p2)

	c1 := Circle{4.56, &Point{1, 2}}
	fmt.Println(c1.x, c1.y)
}
