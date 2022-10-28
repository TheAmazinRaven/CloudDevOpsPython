package main

import (
	"fmt"
)

func main() {
	var grade int
	var letterGrade string

	fmt.Println("What is your grade?")
	fmt.Scan(&grade)

	if grade > 90 {
		letterGrade = "A"
		fmt.Printf("Your grade is %v. Your letter grade is %v.", grade, letterGrade)
	} else if grade > 75 && grade < 89 {
		letterGrade = "B"
		fmt.Printf("Your grade is %v. Your letter grade is %v.", grade, letterGrade)
	} else {
		letterGrade = "C"
		fmt.Printf("Your grade is %v. Your letter grade is %v.", grade, letterGrade)
	}
}
