package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50


	fmt.Printf("Welcome to %s booking application!\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")

	// var userName
	// // Ask for their name
	// userName = "Alice"
	// fmt.Printf(userName)

}