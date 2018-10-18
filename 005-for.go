package main

import "fmt"

func main() {

	//
	fmt.Println("Single loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//
	fmt.Println("Regular loop")
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	//
	fmt.Println("Infinite loop with break")
	for {
		fmt.Println("It breaks...")
		break
	}

	//
	fmt.Println("Loop with continue")
	for k := 0; k <= 8; k++ {
		if k%2 == 1 {
			continue
		}

		fmt.Println(k)
	}
}
