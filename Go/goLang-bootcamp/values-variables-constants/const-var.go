package main

import (
	"fmt"
)

func main() {
	// defined string variables
	var name string
	name = "Raven"

	fmt.Println("My name is", name)

	// reassign variables
	var name2 string
	name2 = "Janet"
	fmt.Println("Hey", name2)

	name2 = "Henry"
	fmt.Println("Hey", name2)

	// another way to define variables
	var name3 = "Raven D"
	fmt.Println("Hey", name3)

	// another another way to declare variables
	name4 := "var new way"
	fmt.Println(name4)

	// different variable types
	game := "XO"         // string
	numberOfPlayers := 2 // int
	gameStarted := false // bool

	fmt.Println("Hi! Welcome to the", game, ". There are", numberOfPlayers, "players! Has your game started?", gameStarted)

	var name5 string
	var number2 int
	var floatNumber float32
	var boo bool

	fmt.Println("String default value:", name5)
	fmt.Println("Number2 default value:", number2)
	fmt.Println("floatNumber default value:", floatNumber)
	fmt.Println("Boo default value:", boo)

	// Constants
	const gameName = "XO"
	fmt.Println(gameName)
}
