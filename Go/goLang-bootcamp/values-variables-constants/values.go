package main

import (
	"fmt"
)

func main() {
	// strings val
	fmt.Println("123 + 123") // 123 + 123

	fmt.Println("hello", "1", "Raven") // Hello 1 Raven

	// int
	fmt.Println(123 + 123) // 246
	fmt.Println(4 / 2)     // 2
	fmt.Println(4 / 5)     // 0

	// string & int
	fmt.Println("Hello", 2, 3)

	// float
	fmt.Println("This is a float number", (7.0 / 2.0), "stuffs")
	fmt.Println("This is NOT a float number", (7 / 2), "stuffs")

	// boolean

	fmt.Println("This is a boolean", true)
	fmt.Println("This is also a boolean", false)

}
