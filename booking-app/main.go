package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "GO Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your GO-tickets here in order to attend")

	for {
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

		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("The GO conference is sold out. We will see you next year. \n")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, You can not book %v tickets\n", remainingTickets, userTickets)
		}
	}
}
