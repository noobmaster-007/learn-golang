package main

import "fmt"

func main() {

	// array default value
	// int - 0
	// bool - false
	// string - empty space
	// float =0
	arr := [3]int{1, 2, 3}

	fmt.Println(arr)

	fmt.Println("Print with for loop")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("Print with range")
	for v := range arr {
		fmt.Println(v)
	}
}
