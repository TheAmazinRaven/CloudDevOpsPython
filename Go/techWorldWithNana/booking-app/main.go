package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, uint(remainingTickets))

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Println(conferenceName)

	for /* remainingTickets > 0 && len(bookings) < 50 */ {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		// Get name from user
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// Check User Validation

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// isValidCity := city == "Singapore" || city == "London"

		if isValidName && isValidEmail && isValidTicketNumber /* userTickets <= remainingTickets */ {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			/*
				fmt.Printf("The whole slice: %v\n", bookings)
				fmt.Printf("The first value: %v\n", bookings[0])
				fmt.Printf("Slice type: %T\n", bookings)
				fmt.Printf("Slice length: %v\n", len(bookings))
			*/

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// call function print first names
			firstNames := getFirstNames(bookings)
			fnt.Printf("The first names of the bookings are: %v\n", firstName)

			if remainingTickets == 0 {
				// ending the program
				fmt.Println("Our conference is booked up! Come back next year!")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("First name or last name you entered is invalid.")
			}

			if !isValidEmail {
				fmt.Println("The email address you entered does not contain an @ symbol.")
			}

			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid.")
			}

			/* fmt.Printf("We only have %v tickets left. So you can't book %v tickets.\n", remainingTickets, userTickets) */
			/* fmt.Printf("Your input is invalid, try again.") */
			continue
		}

	}

}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames(bookings []string []string) {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
