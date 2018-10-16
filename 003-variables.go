package main

import "fmt"

func main() {

	// Automatic types - String
	var hello = "Hello World!"
	fmt.Println(hello)

	// Automatic types - Integer
	var ia, ib = 1, 2
	fmt.Println("integer a =", ia)
	fmt.Println("integer b =", ib)

	// Automatic types- float
	var fa, fb float32 = 1, 2
	fmt.Println("float a =", fa)
	fmt.Println("float b =", fb)

	// Zero value as default
	var i int
	fmt.Println("Integer i =", i)

	// Define without "var"
	str := "Without var"
	fmt.Println(str)
}
