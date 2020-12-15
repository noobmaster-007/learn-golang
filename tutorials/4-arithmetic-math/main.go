package main

import "fmt"

func main() {
	// Arithmetic must be used for the same type

	// division with integer
	var num1 int = 9
	var num2 int = 4

	fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)

	// number need to convert to float if we need get the float result
	fmt.Printf("%.2f / %.2f = %.2f\n", float64(12), float64(5), float64(12)/float64(5))
}
