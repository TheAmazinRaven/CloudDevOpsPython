package main

import "fmt"

func main() {

	printHelloWorld("Raven", 9)
	printHelloWorld("James", 4)
	summer := mathGuy(7, 5)
	println(summer)
	println(mathGuy(8, 2))
	human("Raven", 28)
	coolHuman("Nathan", 12)
	test()
}

func printHelloWorld(something string, thing int) {
	println("Hello world! :)")
	println(something, thing)
}

func mathGuy(x int, y int) int {
	sum := x + y
	return sum
}

func human(name string, age int) {
	println(name, age)
}

func coolHuman(name string, age int) {
	human(name, age)
}

func test() {
	fmt.Println("test")
}
