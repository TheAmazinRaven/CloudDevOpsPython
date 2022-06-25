package main // declaring a main package

import "fmt" // imporing the format function
import "rsc.io/quote"

func main() { // executesby default when code is ran
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
