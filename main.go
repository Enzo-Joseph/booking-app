package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Initialize booking variables
const conferenceName string = "Go Conference"

var remainingTickets uint = 50
var bookings []UserData = make([]UserData, 0)
var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greatUser()

	// Get user bookings
	for remainingTickets > 0 {

		userFirstName, userLastName, userEmail, userBookedTickets := getUserBooking()

		if !helper.IsValidInput(userFirstName, userLastName, userEmail, userBookedTickets, remainingTickets) {
			continue
		}

		// else, the inputs are valid
		bookTicket(userFirstName, userLastName, userEmail, userBookedTickets)

		wg.Add(1)
		go sendTicket(userFirstName, userLastName, userEmail, userBookedTickets)

		var firstNames []string = getFirstNames()
		fmt.Printf("Every booking : %v\n", firstNames)

	}

	fmt.Printf("Every tickets for %v are booked.\n", conferenceName)

	wg.Wait()

}

func greatUser() {
	fmt.Printf("Welcome to %v booking app.\n", conferenceName)
	fmt.Printf("There are %v tickets available.\n", remainingTickets)
}

func getUserBooking() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userBookedTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&userFirstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&userLastName)
	fmt.Println("Enter your email")
	fmt.Scan(&userEmail)
	fmt.Println("Enter the number of tickets you want to book")
	fmt.Scan(&userBookedTickets)

	return userFirstName, userLastName, userEmail, userBookedTickets
}

func getFirstNames() []string {

	var firstNames []string = []string{}
	for _, booking := range bookings {
		var firstName string = booking.firstName
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func bookTicket(userFirstName string, userLastName string, userEmail string, userBookedTickets uint) {
	var userData = UserData{
		firstName:       userFirstName,
		lastName:        userLastName,
		email:           userEmail,
		numberOfTickets: userBookedTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets to %v. You will receive a confirmation at %v\n", userFirstName, userLastName, userBookedTickets, conferenceName, userEmail)

	remainingTickets = remainingTickets - userBookedTickets
	fmt.Printf("There are %v remaining tickets for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(firstName string, lastName string, email string, numberOfTickets uint) {
	time.Sleep(10 * time.Second)
	fmt.Printf("\n#########################\n")
	var ticket = fmt.Sprintf("%v tickets for %v %v.", numberOfTickets, firstName, lastName)
	fmt.Printf("Ticket :\n%v\nSent to email adress %v\n", ticket, email)
	fmt.Printf("#########################\n")
	wg.Done()
}
