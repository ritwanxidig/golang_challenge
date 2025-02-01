package main

import (
	"booking-app/helpers"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	userName        string
	email           string
	numberOfTickets uint
}

func main() {

	helpers.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	// for remainingTickets > 0 && len(bookings) < conferenceTickets {

	var userName, email, userTickets = getUserInput()

	var isValidName, isValidEmail, isValidTicketNumber = helpers.ValidateInputValues(userName, email, userTickets, remainingTickets)

	if !isValidName {
		fmt.Println("Your name is too short. Please try again")
		// continue
	}

	if !isValidEmail {
		fmt.Println("Your email is invalid. Please try again")
		// continue
	}

	if !isValidTicketNumber {
		fmt.Println("Invalid number of tickets. Please try again")
		// continue
	}

	if userTickets < remainingTickets {

		bookTicket(userTickets, userName, email)

		wg.Add(1)
		go sendTicket(userName, email, userTickets)

		fmt.Printf("Thank you %v for buying %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		var firstNames = extractFirstNames()

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else if userTickets == remainingTickets {
		// do something
		fmt.Printf("We have only %v tickets remaining, so we cannot sell all the tickets only to you\n", remainingTickets)
	} else {
		fmt.Printf("Sorry, we have only %v tickets remaining, please try again\n", remainingTickets)
		// continue
		// }
	}
	wg.Wait()
}

func getUserInput() (string, string, uint) { // asking user for their name and email
	userName := ""
	email := ""
	var userTickets uint

	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	for {
		var userTicketInput string
		fmt.Print("How many tickets do you want to buy? ")
		fmt.Scan(&userTicketInput)

		tickets, err := strconv.ParseUint(userTicketInput, 10, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid positive number.")
			continue
		}

		userTickets = uint(tickets)
		break
	}

	return userName, email, userTickets
}

func bookTicket(userTickets uint, userName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var myUser = userData{
		userName:        userName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, myUser)
}

func extractFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var fullName = strings.Fields(booking.userName)
		var firstName = fullName[0]
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}

func sendTicket(userName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	fmt.Println("############### Email Sending ... #####################")
	var ticket = fmt.Sprintf("%v tickets for %v", userTickets, userName)
	fmt.Printf("Sending ticket: \n %v \n to email address %v \n", ticket, email)
	fmt.Println("############### Email Sent ! #####################")
	wg.Done()
}
