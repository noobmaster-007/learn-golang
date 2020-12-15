package main

import "fmt"

// Main function
func main() {

	// UINT8 overflow
	var number uint8 = 255
	number = number + 1
	fmt.Println("Number overflow", number)

	// See what is the type compiler will choose
	flag := true
	integer := 230
	decimal := 12.23
	str := "Name"

	// Explicitly defined type
	var uint8Num uint8 = 239

	fmt.Printf("Type of %v is %T\n", flag, flag)
	fmt.Printf("Type of %v is %T\n", integer, integer)
	fmt.Printf("Type of %v is %T\n", decimal, decimal)
	fmt.Printf("Type of %v is %T\n", str, str)
	fmt.Printf("Type of %v is %T\n", uint8Num, uint8Num)

	// Following code will not work because num will have integer number,
	// but we are try to assign a floating number to it
	// num := 20
	// num = 12.22

}
