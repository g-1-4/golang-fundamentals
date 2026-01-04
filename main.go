package main

import (
	"fmt"
	"strings"
)

var conferenceName string= "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	// conferenceName := "Go Conference"
	// const conferenceTickets int = 50
	// var remainingTickets uint = 50
	// bookings := []string{}

	// var bookings [50]string

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T", conferenceName, conferenceTickets, remainingTickets)

	// fmt.Printf("Welcome to %s booking application!\n", conferenceName)
	// fmt.Printf("We have a total of %d tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your ticket here to attend.")

	greetUsers()


	for {

		// var userName string

		// var firstName string
		// var lastName string
		// var email string
		// var userTickets uint

		// userName = "Alice"
		// userTickets = 2

		// fmt.Printf("User %s is booking %d tickets.\n", userName, userTickets)

		// User Input
		// fmt.Println("Enter your first name: ")
		// fmt.Scan(&firstName)

		// fmt.Println("Enter your last name: ")
		// fmt.Scan(&lastName)

		// fmt.Println("Enter your email address: ")
		// fmt.Scan(&email)

		// fmt.Println("Enter number of tickets: ")
		// fmt.Scan(&userTickets)

		firstName, lastName, email, userTickets := getUserInput()

		// Input validation
		// isValidName := len(firstName)>=2 && len(lastName)>=2
		// isValidEmail := strings.Contains(email,"@")
		// isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)


		if isValidName && isValidEmail && isValidTicketNumber {
			// remainingTickets = remainingTickets - userTickets
			// bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first booking: %v\n", bookings[0])
			// fmt.Printf("Array type: %T\n", bookings)
			// fmt.Printf("Array length: %v\n", len(bookings))

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first booking: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))

			// fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s.\n", firstName, lastName, userTickets, email)
			// fmt.Printf("%d tickets remaining for %s.\n", remainingTickets, conferenceName)

			// firstNames := []string{}
			// for booking := range bookings {
			// 	names := strings.Fields(bookings[booking])
			// 	firstNames = append(firstNames, names[0])
			// }
			// fmt.Printf("These are all our bookings: %v\n", firstNames)

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := []string{}
			firstNames = getFirstNames(firstNames)
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our conference is booked out. Come back next year.\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
			fmt.Println("Your input data is invalid, Please try again.")
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %s booking application!\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")
}

func getFirstNames(firstNames []string) []string {
	for booking := range bookings {
		names := strings.Fields(bookings[booking])
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName)>=2 && len(lastName)>=2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %s.\n", remainingTickets, conferenceName)	
}