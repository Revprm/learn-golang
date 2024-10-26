package main

import "fmt"

func main(){
	// variables and Data Types
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// arrays and slices
	bookings := []string{}
	
	// conferenceTickets = 100 // cannot assign to conferenceTickets as it is a constant
	fmt.Printf("conferenceTickets is %T\nremainingTickets is %T\nconferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println	("Get your tickets here to attend")




	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user to enter their name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your Email Address")
	fmt.Scan(&email)

	fmt.Println("How many tickets you want to book?")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	// bookings[0] = firstName + " " + lastName

	bookings = append(bookings, firstName + " " + lastName)

	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)

	// firstName = "Revy"
	// userTickets = 2

	fmt.Printf("The whole array is %v\n", bookings)
	fmt.Printf("The first value is %v\n", bookings[0])
	fmt.Printf("Slice type is %T\n", bookings)
	fmt.Printf("The Slice length is %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets for %v. You will receive a confirmation Email at %v\n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("We have %v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings is %v\n", bookings)
}
