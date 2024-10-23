package main

import (
	"fmt"
	"strings"
)

func main() {
	// Initialize booking variables
	const conferenceName string = "Go Conference"
	var remainingTickets uint = 50

	greatUser(conferenceName, remainingTickets)

	// Get user bookings
	for remainingTickets > 0 {
		var userName string
		var userLastName string
		var userEmail string
		var userBookedTickets uint
		userName, userLastName, userEmail, userBookedTickets = getUserBooking()

		if !isValidInput(userName, userLastName, userEmail, userBookedTickets, remainingTickets) {
			continue
		}

		// else, the inputs are valid

		fmt.Printf("Thank you %v %v for booking %v tickets to %v. You will receive a confirmation at %v\n", userName, userLastName, userBookedTickets, conferenceName, userEmail)

		remainingTickets = remainingTickets - userBookedTickets
		fmt.Printf("There are %v remaining tickets for %v.\n", remainingTickets, conferenceName)

	}

	fmt.Printf("Every tickets for %v are booked.\n", conferenceName)

}

func greatUser(conferenceName string, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking app.\n", conferenceName)
	fmt.Printf("There are %v tickets available.\n", remainingTickets)
}

func getUserBooking() (string, string, string, uint) {
	var userName string
	var userLastName string
	var userEmail string
	var userBookedTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&userName)
	fmt.Println("Enter your last name")
	fmt.Scan(&userLastName)
	fmt.Println("Enter your email")
	fmt.Scan(&userEmail)
	fmt.Println("Enter the number of tickets you want to book")
	fmt.Scan(&userBookedTickets)

	return userName, userLastName, userEmail, userBookedTickets
}

func isValidInput(userName string, userLastName string, userEmail string, userBookedTickets uint, remainingTickets uint) bool {
	var validNames bool = len(userName) > 0 && len(userLastName) > 0
	var validEmail bool = strings.Contains(userEmail, "@")
	var validBookedTickets bool = userBookedTickets <= remainingTickets

	if !validNames {
		fmt.Println("Invalid name.")
	} else if !validEmail {
		fmt.Println("Invalid email.")
	} else if !validBookedTickets {
		fmt.Println("Invalid number of tickets booked.")
	}

	return validNames && validEmail && validBookedTickets
}
