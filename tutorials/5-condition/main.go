package main

import "fmt"

func main() {

	// Compare
	fmt.Printf("5 < 6: %t\n", 5 < 6)
	fmt.Printf("5 > 6: %t\n", 5 > 6)
	fmt.Printf("5 == 6: %t\n", 5 == 6)
	fmt.Printf("5 > 4.5 = %t\n", 5 == 4.5)

	// Compare with diff types
	x := 5
	y := 4.5
	fmt.Printf("%f > %f = %t\n", float64(x), y, float64(x) > y)

	// string compare if based on ASCI value
	fmt.Printf("%s < %s = %t\n", "a", "b", "a" < "b")
	fmt.Printf("%s < %s = %t\n", "A", "b", "A" < "b")
	fmt.Printf("%s == %s = %t\n", "A", "a", "A" == "a")

	// if, else if, else
	// see 3) input-console-type
}
