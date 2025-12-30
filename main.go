package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	// var bookings [50]string

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %s booking application!\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")


	for {

		// var userName string

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// userName = "Alice"
		// userTickets = 2

		// fmt.Printf("User %s is booking %d tickets.\n", userName, userTickets)

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first booking: %v\n", bookings[0])
			// fmt.Printf("Array type: %T\n", bookings)
			// fmt.Printf("Array length: %v\n", len(bookings))

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first booking: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%d tickets remaining for %s.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our conference is booked out. Come back next year.\n")
				break
			}
		} else {
			fmt.Printf("We only have %d tickets remaining, so you can't book %d tickets\n", remainingTickets, userTickets)
			fmt.Println("Please try again")
		}
	}
}