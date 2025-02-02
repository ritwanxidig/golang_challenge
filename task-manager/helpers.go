package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCreateUserInput() (string, bool) {
	reader := bufio.NewReader(os.Stdin) // Create a new reader
	for {
		fmt.Print("Enter task description: ")
		descriptionInput, _ := reader.ReadString('\n')         // Read the entire line (includes spaces)
		descriptionInput = strings.TrimSpace(descriptionInput) // Trim extra spaces and newline
		var isValidDescription bool = len(descriptionInput) > 4
		if isValidDescription {
			return descriptionInput, false
		}
		fmt.Println("Description must be at least 5 characters long.")
	}
}

func getUserIdInput() int {
	var userIdInput string
	for {
		fmt.Print("Enter user ID (must be a positive number): ")
		fmt.Scanln(&userIdInput)
		userId, err := strconv.Atoi(userIdInput)
		if err != nil || userId <= 0 {
			fmt.Println("Invalid user ID, please try again.")
			continue
		}
		return userId
	}
}

func getConfirmInput() int {
	var option string
	for {
		fmt.Print("Please confirm? (yes/no): ")
		fmt.Printf("\n 1. yes\n 2. no\n")
		fmt.Scanln(&option)

		var convertedValue, err = strconv.Atoi(option)
		if err == nil && (convertedValue == 1 || convertedValue == 2) {
			return convertedValue
		}
		fmt.Println("Invalid option. Please enter either 1 or 2.")
		continue
	}
}
