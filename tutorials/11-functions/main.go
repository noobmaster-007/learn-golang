package main

import "fmt"

func main() {

	fmt.Printf("abs(-10) = %d\n", abs(-10))
	fmt.Printf("1 + 1 = %d\n", add(1, 1))
	fmt.Println(multi(1, 2))
	fmt.Println(multi2(1, 2))

	// assign function to a variable
	x := abs
	fmt.Println(x(-12))

	// create a function inside function
	// call the function inline
	test := func(x int) int {
		return x * -1
	}

	fmt.Println(test(10))

	test2(test)

	returnFunc("Hello")()
}

// Pass function into another function
func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

// return function
func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

// syntax of function
// func funcName (param paramType, ...) returnType
func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}

// Param with same type
func add(x, y int) int {
	return x + y
}

// multiple return types
func multi(x, y int) (int, int) {
	return x + y, x - y
}

// multiple return with name
// defer will add statement to stack
func multi2(x, y int) (n1 int, n2 int) {
	// defer will run before return
	defer fmt.Println("Hello")
	n1 = x + y
	n2 = x - y
	defer fmt.Println("Before return")
	return
}
