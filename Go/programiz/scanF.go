package main

import (
	"fmt"
)

func main() {
	var name, color string
	fmt.Println("What is your name and favorite color?")
	fmt.Scan(&name, &color)

	fmt.Printf("Your name is %s, and your favorite color is %s.", name, color)
}
