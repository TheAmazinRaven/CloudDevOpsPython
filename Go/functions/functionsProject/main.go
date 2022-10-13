package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the simple game!")
	name()
	age()
	color()
}

func name(name string) {
	fmt.Println("What is your name?")
	fmt.Scanln(&name)
}

func age(age int) {
	fmt.Println("")
}

func color(color string) {
	fmt.Println("")
}
