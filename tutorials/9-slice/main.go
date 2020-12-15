package main

import "fmt"

func main() {
	// x will be created as an array as size
	// is provided
	x := [5]int{1, 2, 3, 4, 5}

	// without size provide, a slice will be created
	// slice index is inclusive begin and exclude the end
	var s []int = x[0:2]
	fmt.Println(s)
	fmt.Printf("Len of s is %d\n", len(s))
	fmt.Printf("Cap of s is %d\n", cap(s))

	// slice is a pointer to the underlayer array
	fmt.Println(s[1:])

	// make a slice
	var a []int = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Len of s is %d\n", len(a))
	fmt.Printf("Cap of s is %d\n", cap(a))
	// cap of slice reduce
	fmt.Printf("Cap of slice of slice %d\n", cap(a[2:3]))

	// append function
	b := append(a, 8, 9)
	fmt.Println(b)
	fmt.Printf("Type is %T\n", b)

	// using make to create slice
	c := make([]int, 5)
	fmt.Println(c)
	fmt.Printf("Type is %T\n", c)
}
