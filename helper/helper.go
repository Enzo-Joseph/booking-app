package helper

import (
	"fmt"
	"strings"
)

func IsValidInput(userName string, userLastName string, userEmail string, userBookedTickets uint, remainingTickets uint) bool {
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
