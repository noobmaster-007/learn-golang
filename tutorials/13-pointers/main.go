package main

import "fmt"

func changeValue(str *string) {
	*str = "changed"
}

func changeValue2(str string) {
	str = "changed2"
}

func main() {
	// & get the pointer or address
	// * with vaue -  de-referene, get the value from address
	// * with data type - define a pointer

	x := 7
	y := &x
	fmt.Println(x, y)
	*y = 8
	fmt.Println(x, y)

	toChange := "Hello"
	changeValue(&toChange)
	fmt.Println(toChange)
}
