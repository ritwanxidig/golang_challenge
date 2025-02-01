package helpers

import (
	"fmt"
	"strings"
)

func GreetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v conference \n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here")
}

func ValidateInputValues(userName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
