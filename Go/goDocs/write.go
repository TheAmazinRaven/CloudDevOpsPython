package main // declaring a main package

import (
	"fmt"

	"rsc.io/quote"
) // importing the format function

func main() { // executes by default when code is ran
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
