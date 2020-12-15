package main

import "fmt"

func main() {

	// Standard for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum of 0-10 = %d\n", sum)

	// while loop
	num := 0
	idx := 0
	for idx < 10 {
		num += idx
		idx++
	}
	fmt.Printf("sum of 0-10 = %d\n", num)

	// infinete loop
	// break - exit loop
	// continue - skip to beginning
	counter := 0
	s := 0
	for {
		if counter%2 == 0 {
			counter++
			continue
		}
		if counter > 10 {
			break
		}
		s += counter
		counter++
	}

	fmt.Printf("sum of odds 0-10 = %d\n", s)
}
