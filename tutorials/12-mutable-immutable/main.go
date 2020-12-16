package main

import "fmt"

func main() {

	// immutable
	// int, float, int, bool, array

	// int
	var x int = 4
	y := x
	x = 12
	fmt.Println(x, y)

	// array
	var arrayX [2]int = [2]int{1, 2}
	arrayY := arrayX
	arrayX[0] = 0
	fmt.Println(arrayX, arrayY)

	// mutable
	// slice
	var sliceX []int = []int{1, 2, 3}
	sliceY := sliceX
	sliceY[0] = 0
	fmt.Println(sliceX, sliceY)

	// map
	var mX map[string]int = map[string]int{"12": 12, "3": 3}
	mY := mX
	mY["4"] = 4
	fmt.Println(mX, mY)

	var changeSlice []int = []int{2, 3, 4}
	changeFirst(changeSlice)
	fmt.Println(changeSlice)

	sx, sy := 1, 10
	sx, sy = swap(sx, sy)
	fmt.Println(sx, sy)

}

func changeFirst(slice []int) {
	slice[0] = 1000
}

func swap(x, y int) (int, int) {
	return y, x
}
