package main

import "fmt"

func main() {
	// create a map
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   9,
		"orange": 8,
	}
	fmt.Println(mp)
	mp["apple"] = 10
	fmt.Println(mp["apple"])

	// delete item from item
	delete(mp, "apple")
	fmt.Println(mp)

	// check whether key exists
	val, ok := mp["apple"]
	if !ok {
		fmt.Println("apple not exists in mp")
	} else {
		fmt.Println(val)
	}

	// create a empty map
	m := make(map[string]int)
	m["age"] = 20
	fmt.Println(m)
}
