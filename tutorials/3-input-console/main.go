package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your score: ")
	scanner.Scan()
	score, err := strconv.ParseUint(scanner.Text(), 10, 8)
	if err != nil {
		fmt.Printf("Getting error %v\n", err)
	} else {
		fmt.Printf("You have score %v with rate rate is: %s\n", score, rate(score))
	}
}

func rate(score uint64) string {

	if score >= 85 {
		return "HD"
	} else if score >= 75 {
		return "D"
	} else if score >= 65 {
		return "C"
	} else if score >= 50 {
		return "P"
	} else {
		return "F"
	}
}
