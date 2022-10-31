package main

import (
	"fmt"
)

func main() {
	// if conditional
	name := "Sara"

	if name == "Sara" {
		fmt.Println("Your name is", name, ".")
	}

	// else statements
	name2 := "Kyle"

	if name2 == "Kyle" {
		fmt.Println("Your name is", name2, ".")
	} else {
		fmt.Println("Your name is not", name2, ".")
	}

	// else if

	playerName := "Jack"

	if playerName == "Raven" {
		fmt.Println("Hi Raven!")
	} else if playerName == "Jack" {
		fmt.Println("Hi Jack!")
	} else {
		fmt.Println("Who are you?")
	}

}
