package main

import "fmt"

func main() {
	speed := 100
	force := 2.5

	speed = speed * int(force)

	fmt.Println(speed)
	fmt.Println(force, int(force))

	speed = int(float64(speed) * force)
	fmt.Println(speed)
}
