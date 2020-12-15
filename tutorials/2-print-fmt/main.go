package main

import "fmt"

func main() {
	// Integer
	fmt.Printf("Integer: %b\n", 9623232)
	fmt.Printf("Integer: %d\n", 9623232)
	fmt.Printf("Integer: %o\n", 9623232)
	fmt.Printf("Integer: %x\n", 9623232)

	// Float number
	fmt.Printf("Number: %e\n", 12.42323286454)
	fmt.Printf("Number: %f\n", 12.42323286454)
	fmt.Printf("Number: %g\n", 12.42323286454)

	// String
	fmt.Printf("Hello %s\n", "World")
	fmt.Printf("Hello %q\n", "World")

	// Width & Precision
	fmt.Printf("Number: %f\n", 3.1415926)
	fmt.Printf("Number: %.2f\n", 3.1415926)
	fmt.Printf("Number: %3.2f\n", 3.1415926)

	// paddings
	fmt.Printf("Pads number %09d is my receipt\n", 12)
	fmt.Printf("Pads number %-9d is my reciept\n", 12)

}
