package main

import (
	"fmt"
)

// When to use short vs long variable declaration

/*
	===== Long Variable Declaration =====
	1. If you don't know the initial value
	2. When you need a package scoped variable
	3. When you want to group variables together for greater readability

	===== Short Variable Declaration =====
	1. Most commonly used in Go; use when you know the value of the variable
	2. To keep the code concise and easy to read
	3. For re-declaration
*/

func main() {
	// When you don't know the value of the variable, use long declaration
	var score int // score already = 0
	fmt.Println(score)

	// When you need a package scoped variable
	var version string
	fmt.Println(version)

	// When you need to group together for better readability

	var (
		// related:
		video string

		// closely related
		duration int
		current  int
	)

	fmt.Println(video, duration, current)

	// ==============================================

	// var width, height = 100, 50 - Don't do this.
	width, height := 100, 50

	fmt.Println(width, height)

	width, color := 50, "red"

	fmt.Println(color)
}
