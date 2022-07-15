package main

import (
	"fmt"

	"github.com/TheAmazinRaven/CloudDevOpsPython/Go/goDocs/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
