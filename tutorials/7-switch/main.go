package main

import "fmt"

func main() {
	ans := 20

	switch ans {

	// case can only be the same type
	// multiple matches
	case 1, 2, 3:
		fmt.Println("Under 1,2,3")
	case 4:
		fmt.Println("It is 4")
	default:
		fmt.Println("Not match")
	}
}
