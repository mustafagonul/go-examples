package main

import "fmt"

func main() {

	// Simple if
	if 8%2 == 0 {
		fmt.Println("8 is divisible by 2")
	}

	// If with else
	if 9%2 == 1 {
		fmt.Println("9 is odd")
	} else {
		fmt.Println("9 is even")
	}

	// If with initialization
	if num := 5; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
