package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Quiz Game! \n What is your name?")

	fmt.Printf("Please enter your name: \n")

	var name string

	fmt.Scan(&name)

	fmt.Printf("Hello %v, I hope you are ready to play!", name)

}
